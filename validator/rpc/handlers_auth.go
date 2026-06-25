package rpc

import (
	"net/http"

	"github.com/sila-chain/Sila-Consensus-Core/v7/io/file"
	"github.com/sila-chain/Sila-Consensus-Core/v7/monitoring/tracing/trace"
	"github.com/sila-chain/Sila-Consensus-Core/v7/network/httputil"
	"github.com/sila-chain/Sila-Consensus-Core/v7/validator/accounts/wallet"
	"github.com/pkg/errors"
)

// Initialize returns metadata regarding whether the caller has authenticated and has a wallet.
func (s *Server) Initialize(w http.ResponseWriter, r *http.Request) {
	_, span := trace.StartSpan(r.Context(), "validator.web.Initialize")
	defer span.End()
	walletExists, err := wallet.Exists(s.walletDir)
	if err != nil {
		httputil.HandleError(w, errors.Wrap(err, "Could not check if wallet exists").Error(), http.StatusInternalServerError)
		return
	}
	exists, err := file.Exists(s.authTokenPath, file.Regular)
	if err != nil {
		httputil.HandleError(w, errors.Wrap(err, "Could not check if auth token exists").Error(), http.StatusInternalServerError)
		return
	}
	httputil.WriteJson(w, &InitializeAuthResponse{
		HasSignedUp: exists,
		HasWallet:   walletExists,
	})
}
