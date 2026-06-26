package beacon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sila-chain/Sila-Consensus-Core/v7/api"
	"github.com/sila-chain/Sila-Consensus-Core/v7/api/server/structs"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/blockchain/kzg"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/gloas"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/core/peerdas"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/db"
	"github.com/sila-chain/Sila-Consensus-Core/v7/beacon-chain/rpc/silaapi/shared"
	fieldparams "github.com/sila-chain/Sila-Consensus-Core/v7/config/fieldparams"
	consensusblocks "github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/blocks"
	"github.com/sila-chain/Sila-Consensus-Core/v7/consensus-types/primitives"
	"github.com/sila-chain/Sila-Consensus-Core/v7/encoding/bytesutil"
	"github.com/sila-chain/Sila-Consensus-Core/v7/monitoring/tracing/trace"
	"github.com/sila-chain/Sila-Consensus-Core/v7/network/httputil"
	eth "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
	"github.com/sila-chain/Sila-Consensus-Core/v7/runtime/version"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetSilaPayloadEnvelope retrieves a full sila payload envelope by beacon block root.
// The blinded envelope is fetched from the DB and the full sila payload is reconstructed
// from the EL via sila_getBlockByHash.
func (s *Server) GetSilaPayloadEnvelope(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "beacon.GetSilaPayloadEnvelope")
	defer span.End()

	blockID := r.PathValue("block_id")
	if blockID == "" {
		httputil.HandleError(w, "block_id is required in URL params", http.StatusBadRequest)
		return
	}

	root, err := s.Blocker.BlockRoot(ctx, []byte(blockID))
	if !shared.WriteBlockRootFetchError(w, err) {
		return
	}

	blinded, err := s.BeaconDB.SilaPayloadEnvelope(ctx, root)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			httputil.HandleError(w, "sila payload envelope not found", http.StatusNotFound)
			return
		}
		httputil.HandleError(w, "could not retrieve sila payload envelope: "+err.Error(), http.StatusInternalServerError)
		return
	}
	full, err := s.ExecutionReconstructor.ReconstructSilaPayloadEnvelope(ctx, blinded)
	if err != nil {
		httputil.HandleError(w, "could not reconstruct sila payload envelope: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(api.VersionHeader, version.String(version.Gloas))

	if httputil.RespondWithSsz(r) {
		sszBytes, err := full.MarshalSSZ()
		if err != nil {
			httputil.HandleError(w, "could not marshal envelope to SSZ: "+err.Error(), http.StatusInternalServerError)
			return
		}
		httputil.WriteSsz(w, sszBytes)
		return
	}

	isOptimistic, err := s.OptimisticModeFetcher.IsOptimisticForRoot(ctx, root)
	if err != nil {
		httputil.HandleError(w, "could not check optimistic status: "+err.Error(), http.StatusInternalServerError)
		return
	}
	finalized := s.FinalizationFetcher.IsFinalized(ctx, root)

	jsonEnvelope, err := structs.SignedSilaPayloadEnvelopeFromConsensus(full)
	if err != nil {
		httputil.HandleError(w, "could not convert envelope to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	httputil.WriteJson(w, &structs.GetSilaPayloadEnvelopeResponse{
		Version:             version.String(version.Gloas),
		SilaOptimistic: isOptimistic,
		Finalized:           finalized,
		Data:                jsonEnvelope,
	})
}

// PublishSilaPayloadEnvelope broadcasts a signed envelope. Sila-Payload-Blinded
// selects the body: true=blinded (stateful, BN reconstructs from cache), false=contents (stateless).
// Endpoint: POST /sila/v1/beacon/sila_payload_envelopes
func (s *Server) PublishSilaPayloadEnvelope(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "beacon.PublishSilaPayloadEnvelope")
	defer span.End()
	versionHeader := r.Header.Get(api.VersionHeader)
	if versionHeader == "" {
		httputil.HandleError(w, api.VersionHeader+" header is required", http.StatusBadRequest)
		return
	}
	v, err := version.FromString(versionHeader)
	if err != nil || v < version.Gloas {
		httputil.HandleError(w, api.VersionHeader+" header must be gloas or later", http.StatusBadRequest)
		return
	}
	blindedHeader := r.Header.Get(api.SilaPayloadBlindedHeader)
	if blindedHeader == "" {
		httputil.HandleError(w, api.SilaPayloadBlindedHeader+" header is required", http.StatusBadRequest)
		return
	}
	isBlinded, err := strconv.ParseBool(blindedHeader)
	if err != nil {
		httputil.HandleError(w, "invalid "+api.SilaPayloadBlindedHeader+" value: "+err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		httputil.HandleError(w, "could not read request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if isBlinded {
		s.publishBlindedEnvelope(ctx, w, r, body)
		return
	}
	s.publishEnvelopeContents(ctx, w, r, body)
}

// publishBlindedEnvelope reconstructs the full envelope from cache by beacon_block_root.
// HTR(blinded) == HTR(full), so the validator signature stays valid.
func (s *Server) publishBlindedEnvelope(ctx context.Context, w http.ResponseWriter, r *http.Request, body []byte) {
	signedBlinded := &eth.SignedWireBlindedSilaPayloadEnvelope{}
	if httputil.IsRequestSsz(r) {
		if err := signedBlinded.UnmarshalSSZ(body); err != nil {
			httputil.HandleError(w, "could not decode SSZ blinded envelope: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		var jsonBlinded structs.SignedBlindedSilaPayloadEnvelope
		if err := json.Unmarshal(body, &jsonBlinded); err != nil {
			httputil.HandleError(w, "could not decode JSON blinded envelope: "+err.Error(), http.StatusBadRequest)
			return
		}
		consensus, err := jsonBlinded.ToConsensus()
		if err != nil {
			httputil.HandleError(w, "invalid signed blinded envelope: "+err.Error(), http.StatusBadRequest)
			return
		}
		signedBlinded = consensus
	}
	if signedBlinded.Message == nil {
		httputil.HandleError(w, "blinded envelope message is nil", http.StatusBadRequest)
		return
	}

	cached, ok := s.SilaPayloadEnvelopeCache.Contents()
	if !ok || cached.Envelope == nil {
		httputil.HandleError(w, "no cached sila payload envelope to reconstruct from", http.StatusBadRequest)
		return
	}
	if !bytes.Equal(cached.Envelope.BeaconBlockRoot, signedBlinded.Message.BeaconBlockRoot) {
		httputil.HandleError(w, "cached envelope beacon_block_root does not match blinded envelope", http.StatusBadRequest)
		return
	}
	blindedRoot, err := signedBlinded.Message.HashTreeRoot()
	if err != nil {
		httputil.HandleError(w, "could not hash blinded envelope: "+err.Error(), http.StatusInternalServerError)
		return
	}
	cachedRoot, err := cached.Envelope.HashTreeRoot()
	if err != nil {
		httputil.HandleError(w, "could not hash cached envelope: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if blindedRoot != cachedRoot {
		httputil.HandleError(w, "cached envelope hash tree root does not match blinded envelope", http.StatusBadRequest)
		return
	}

	full := &eth.SignedSilaPayloadEnvelope{
		Message:   cached.Envelope,
		Signature: signedBlinded.Signature,
	}

	if !s.validateEnvelopeBroadcast(ctx, w, r, full) {
		return
	}

	if _, err := s.V1Alpha1ValidatorServer.PublishSilaPayloadEnvelope(ctx, full); err != nil {
		writeEnvelopePublishError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// writeEnvelopePublishError maps the v1alpha1 publish outcome to the spec status
// codes: InvalidArgument -> 400, Aborted -> 202 (broadcast ok, import failed).
func writeEnvelopePublishError(w http.ResponseWriter, err error) {
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.InvalidArgument:
			httputil.HandleError(w, st.Message(), http.StatusBadRequest)
		case codes.Aborted:
			httputil.HandleError(w, st.Message(), http.StatusAccepted)
		default:
			httputil.HandleError(w, st.Message(), http.StatusInternalServerError)
		}
		return
	}
	httputil.HandleError(w, "could not publish sila payload envelope: "+err.Error(), http.StatusInternalServerError)
}

// publishEnvelopeContents handles the stateless flow (header=false).
func (s *Server) publishEnvelopeContents(ctx context.Context, w http.ResponseWriter, r *http.Request, body []byte) {
	if httputil.IsRequestSsz(r) {
		contents := &eth.SignedSilaPayloadEnvelopeContents{}
		if err := contents.UnmarshalSSZ(body); err != nil {
			httputil.HandleError(w, "could not decode SSZ envelope contents: "+err.Error(), http.StatusBadRequest)
			return
		}
		s.publishSilaPayloadEnvelopeContentsSSZ(ctx, w, r, contents)
		return
	}
	s.publishSilaPayloadEnvelopeContents(ctx, w, r, body)
}

// publishSilaPayloadEnvelopeContents handles the JSON stateless variant.
func (s *Server) publishSilaPayloadEnvelopeContents(ctx context.Context, w http.ResponseWriter, r *http.Request, body []byte) {
	var contents structs.SignedSilaPayloadEnvelopeContents
	if err := json.Unmarshal(body, &contents); err != nil {
		httputil.HandleError(w, "could not decode envelope contents: "+err.Error(), http.StatusBadRequest)
		return
	}
	signed, kzgProofs, blobs, err := contents.ToConsensus()
	if err != nil {
		httputil.HandleError(w, "invalid signed sila payload envelope contents: "+err.Error(), http.StatusBadRequest)
		return
	}
	s.processEnvelopeContents(ctx, w, r, signed, kzgProofs, blobs)
}

// publishSilaPayloadEnvelopeContentsSSZ handles the SSZ stateless variant.
func (s *Server) publishSilaPayloadEnvelopeContentsSSZ(ctx context.Context, w http.ResponseWriter, r *http.Request, contents *eth.SignedSilaPayloadEnvelopeContents) {
	if contents == nil || contents.SignedSilaPayloadEnvelope == nil {
		httputil.HandleError(w, "nil signed sila payload envelope contents", http.StatusBadRequest)
		return
	}
	s.processEnvelopeContents(ctx, w, r, contents.SignedSilaPayloadEnvelope, contents.KzgProofs, contents.Blobs)
}

// processEnvelopeContents verifies caller-supplied blobs/proofs, broadcasts
// derived sidecars, then delegates the envelope to the bare publish path.
func (s *Server) processEnvelopeContents(ctx context.Context, w http.ResponseWriter, r *http.Request, signed *eth.SignedSilaPayloadEnvelope, kzgProofs, blobs [][]byte) {
	if !s.validateEnvelopeBroadcast(ctx, w, r, signed) {
		return
	}

	if len(blobs) > 0 {
		blockRoot := bytesutil.ToBytes32(signed.Message.BeaconBlockRoot)
		cellsPerBlob, proofsPerBlob, err := peerdas.ComputeCellsAndProofsFromFlat(blobs, kzgProofs)
		if err != nil {
			httputil.HandleError(w, "could not compute cells and proofs: "+err.Error(), http.StatusBadRequest)
			return
		}
		// External trust boundary — verify before broadcasting/storing.
		if err := verifyCellProofs(blobs, kzgProofs); err != nil {
			httputil.HandleError(w, "kzg verification failed: "+err.Error(), http.StatusBadRequest)
			return
		}
		roSidecars, err := peerdas.DataColumnSidecarsGloas(cellsPerBlob, proofsPerBlob, primitives.Slot(signed.Message.Payload.SlotNumber), blockRoot)
		if err != nil {
			httputil.HandleError(w, "could not build data column sidecars: "+err.Error(), http.StatusInternalServerError)
			return
		}
		verifiedSidecars := make([]consensusblocks.VerifiedRODataColumn, 0, len(roSidecars))
		for _, sc := range roSidecars {
			verifiedSidecars = append(verifiedSidecars, consensusblocks.NewVerifiedRODataColumn(sc))
		}
		if err := s.Broadcaster.BroadcastDataColumnSidecars(ctx, verifiedSidecars); err != nil {
			httputil.HandleError(w, "could not broadcast data column sidecars: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if err := s.DataColumnReceiver.ReceiveDataColumns(verifiedSidecars); err != nil {
			httputil.HandleError(w, "could not receive data column sidecars: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if _, err := s.V1Alpha1ValidatorServer.PublishSilaPayloadEnvelope(ctx, signed); err != nil {
		writeEnvelopePublishError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// validateEnvelopeBroadcast applies broadcast_validation semantics to an
// envelope publish before it is broadcast to gossip. Spec: beacon-APIs #580.
// Writes the HTTP error and returns false on failure: 400 for validation
// failures, 500 for internal errors.
//   - gossip (default): no extra REST-layer checks — the downstream gossip
//     pipeline performs validation.
//   - consensus: full envelope consensus checks against the head state. Submission
//     path requires envRoot to equal head.
//   - consensus_and_equivocation: consensus + reject if a different beacon
//     block at the envelope's slot has already been received.
func (s *Server) validateEnvelopeBroadcast(ctx context.Context, w http.ResponseWriter, r *http.Request, signed *eth.SignedSilaPayloadEnvelope) bool {
	level := r.URL.Query().Get(broadcastValidationQueryParam)
	switch level {
	case "", broadcastValidationGossip:
		// TODO: run lightweight gossip checks (sig + bid consistency) here — beacon-APIs #580.
		return true
	case broadcastValidationConsensus, broadcastValidationConsensusAndEquivocation:
	default:
		httputil.HandleError(w, fmt.Sprintf("invalid %s value: %q", broadcastValidationQueryParam, level), http.StatusBadRequest)
		return false
	}

	envSlot := signed.Message.Payload.SlotNumber
	envRoot := bytesutil.ToBytes32(signed.Message.BeaconBlockRoot)

	if level == broadcastValidationConsensusAndEquivocation {
		// CanonicalNodeAtSlot's bool means "payload full", not "node found" — at the
		// wall clock slot it is always false. A non-zero root is the found signal.
		canonRoot, _ := s.ForkchoiceFetcher.CanonicalNodeAtSlot(envSlot)
		if canonRoot != ([32]byte{}) && canonRoot != envRoot {
			err := errors.Wrapf(errEquivocatedBlock, "another block for slot %d already exists in fork choice", envSlot)
			httputil.HandleError(w, err.Error(), http.StatusBadRequest)
			return false
		}
	}

	// Submission path: envelope must be for the current head.
	headRoot, err := s.HeadFetcher.HeadRoot(ctx)
	if err != nil {
		httputil.HandleError(w, "could not get head root: "+err.Error(), http.StatusInternalServerError)
		return false
	}
	if !bytes.Equal(headRoot, envRoot[:]) {
		httputil.HandleError(w, fmt.Sprintf("envelope beacon block root %#x is not canonical head", envRoot), http.StatusBadRequest)
		return false
	}
	st, err := s.HeadFetcher.HeadState(ctx)
	if err != nil {
		httputil.HandleError(w, "could not get head state: "+err.Error(), http.StatusInternalServerError)
		return false
	}
	roSigned, err := consensusblocks.WrappedROSignedSilaPayloadEnvelope(signed)
	if err != nil {
		httputil.HandleError(w, "could not wrap signed envelope: "+err.Error(), http.StatusInternalServerError)
		return false
	}
	if err := gloas.VerifySilaPayloadEnvelope(ctx, st, roSigned); err != nil {
		httputil.HandleError(w, "consensus validation failed: "+err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

// verifyCellProofs batch-verifies cell proofs against commitments derived
// from the supplied blobs. Does not tie data to a specific block — that needs
// the block's BlobKzgCommitments which a stateless receiver may not have.
func verifyCellProofs(blobs [][]byte, flatProofs [][]byte) error {
	commitments := make([][]byte, len(blobs))
	for i, blob := range blobs {
		if len(blob) != len(kzg.Blob{}) {
			return errors.Errorf("blob %d has wrong size %d", i, len(blob))
		}
		var b kzg.Blob
		copy(b[:], blob)
		c, err := kzg.BlobToKZGCommitment(&b)
		if err != nil {
			return errors.Wrapf(err, "compute kzg commitment for blob %d", i)
		}
		commitments[i] = c[:]
	}
	return kzg.VerifyCellKZGProofBatchFromBlobData(blobs, commitments, flatProofs, fieldparams.NumberOfColumns)
}

// PublishSignedSilaPayloadBid broadcasts a signed sila payload bid to the P2P network.
func (s *Server) PublishSignedSilaPayloadBid(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "beacon.PublishSignedSilaPayloadBid")
	defer span.End()

	if shared.IsSyncing(ctx, w, s.SyncChecker, s.HeadFetcher, s.TimeFetcher, s.OptimisticModeFetcher) {
		return
	}

	versionHeader := r.Header.Get(api.VersionHeader)
	if versionHeader == "" {
		httputil.HandleError(w, api.VersionHeader+" header is required", http.StatusBadRequest)
		return
	}

	var signedBid *eth.SignedSilaPayloadBid
	if httputil.IsRequestSsz(r) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			httputil.HandleError(w, "Could not read request body: "+err.Error(), http.StatusBadRequest)
			return
		}
		signedBid = &eth.SignedSilaPayloadBid{}
		if err := signedBid.UnmarshalSSZ(body); err != nil {
			httputil.HandleError(w, "Could not unmarshal SSZ: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		var jsonBid structs.SignedSilaPayloadBid
		if err := json.NewDecoder(r.Body).Decode(&jsonBid); err != nil {
			if errors.Is(err, io.EOF) {
				httputil.HandleError(w, "No data submitted", http.StatusBadRequest)
				return
			}
			httputil.HandleError(w, "Could not decode request body: "+err.Error(), http.StatusBadRequest)
			return
		}
		var err error
		signedBid, err = jsonBid.ToConsensus()
		if err != nil {
			httputil.HandleError(w, "Could not convert bid to consensus type: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := s.Broadcaster.Broadcast(ctx, signedBid); err != nil {
		httputil.HandleError(w, "Could not broadcast sila payload bid: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
