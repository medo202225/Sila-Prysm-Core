// Package params defines all custom parameter configurations
// for running end to end tests.
package params

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/sila-chain/Sila-Consensus-Core/v7/io/file"
	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/sila-chain/Sila/core/types"
)

// params struct defines the parameters needed for running E2E tests to properly handle test sharding.
type params struct {
	TestPath                  string
	LogPath                   string
	TestShardIndex            int
	BeaconNodeCount           int
	LighthouseBeaconNodeCount int
	Ports                     *ports
	Paths                     *paths
	SilaExecutionGenesisBlock          *types.Block
	StartTime                 time.Time
	CLGenesisTime             time.Time
	SilaExecutionGenesisTime           time.Time
	NumberOfExecutionCreds    uint64
}

type ports struct {
	BootNodePort                    int
	BootNodeMetricsPort             int
	SilaExecutionPort                        int
	SilaExecutionRPCPort                     int
	SilaExecutionAuthRPCPort                 int
	SilaExecutionWSPort                      int
	SilaExecutionProxyPort                   int
	SilaBeaconNodeRPCPort          int
	SilaBeaconNodeUDPPort          int
	SilaBeaconNodeQUICPort         int
	SilaBeaconNodeTCPPort          int
	SilaBeaconNodeHTTPPort         int
	SilaBeaconNodeMetricsPort      int
	SilaBeaconNodePprofPort        int
	LighthouseBeaconNodeP2PPort     int
	LighthouseBeaconNodeHTTPPort    int
	LighthouseBeaconNodeMetricsPort int
	ValidatorMetricsPort            int
	ValidatorHTTPPort               int
	JaegerTracingPort               int
}

type paths struct{}

// SilaExecutionStaticFile abstracts the location of the silaexec static file folder in the e2e directory, so that
// a relative path can be used.
// The relative path is specified as a variadic slice of path parts, in the same way as path.Join.
func (*paths) SilaExecutionStaticFile(rel ...string) string {
	parts := append([]string{SilaExecutionStaticFilesPath}, rel...)
	return path.Join(parts...)
}

// SilaExecutionRunfile returns the full path to a file in the silaexec static directory, within bazel's run context.
// The relative path is specified as a variadic slice of path parts, in the same style as path.Join.
func (p *paths) SilaExecutionRunfile(rel ...string) (string, error) {
	return bazel.Runfile(p.SilaExecutionStaticFile(rel...))
}

// MinerKeyPath returns the full path to the file containing the miner's cryptographic keys.
func (p *paths) MinerKeyPath() (string, error) {
	return p.SilaExecutionRunfile(minerKeyFilename)
}

// TestParams is the globally accessible var for getting config elements.
var TestParams *params

// Logfile gives the full path to a file in the bazel test environment log directory.
// The relative path is specified as a variadic slice of path parts, in the same style as path.Join.
func (p *params) Logfile(rel ...string) string {
	return path.Join(append([]string{p.LogPath}, rel...)...)
}

// SilaExecutionRPCURL gives the full url to use to connect to the given silaexec client's RPC endpoint.
// The `index` param corresponds to the `index` field of the `silaexec.Node` e2e component.
// These are off by one compared to corresponding beacon nodes, because the miner is assigned index 0.
// eg instance the index of the EL instance associated with beacon node index `0` would typically be `1`.
func (p *params) SilaExecutionRPCURL(index int) *url.URL {
	return &url.URL{
		Scheme: baseELScheme,
		Host:   net.JoinHostPort(baseELHost, fmt.Sprintf("%d", p.Ports.SilaExecutionRPCPort+index)),
	}
}

// BootNodeLogFileName is the file name used for the beacon chain node logs.
var BootNodeLogFileName = "bootnode.log"

// TracingRequestSinkFileName is the file name for writing raw trace requests.
var TracingRequestSinkFileName = "tracing-http-requests.log.gz"

// BeaconNodeLogFileName is the file name used for the beacon chain node logs.
var BeaconNodeLogFileName = "beacon-%d.log"

// ValidatorLogFileName is the file name used for the validator client logs.
var ValidatorLogFileName = "vals-%d.log"

// StandardBeaconCount is a global constant for the count of beacon nodes of standard E2E tests.
var StandardBeaconCount = 2

// StandardLighthouseNodeCount is a global constant for the count of lighthouse beacon nodes of standard E2E tests.
var StandardLighthouseNodeCount = 2

// DepositCount is the number of deposits the E2E runner should make to evaluate post-genesis deposit processing.
var DepositCount = uint64(64)

// PostElectraDepositCount is the number of deposits the E2E runner should make to evaluate post-electra deposit processing.
var PostElectraDepositCount = uint64(32)

// PregenesisExecCreds is the number of withdrawal credentials of genesis validators which use an execution address.
var PregenesisExecCreds = uint64(8)

// Base port values.
const (
	portSpan = 50

	bootNodePort        = 2150
	bootNodeMetricsPort = bootNodePort + portSpan

	silaexecPort        = 3150
	silaexecRPCPort     = silaexecPort + portSpan
	silaexecWSPort      = silaexecPort + 2*portSpan
	silaexecAuthRPCPort = silaexecPort + 3*portSpan
	silaexecProxyPort   = silaexecPort + 4*portSpan

	silaBeaconNodeRPCPort     = 4150
	silaBeaconNodeUDPPort     = silaBeaconNodeRPCPort + portSpan
	silaBeaconNodeQUICPort    = silaBeaconNodeRPCPort + 2*portSpan
	silaBeaconNodeTCPPort     = silaBeaconNodeRPCPort + 3*portSpan
	silaBeaconNodeHTTPPort    = silaBeaconNodeRPCPort + 4*portSpan
	silaBeaconNodeMetricsPort = silaBeaconNodeRPCPort + 5*portSpan
	silaBeaconNodePprofPort   = silaBeaconNodeRPCPort + 6*portSpan

	lighthouseBeaconNodeP2PPort     = 5150
	lighthouseBeaconNodeHTTPPort    = lighthouseBeaconNodeP2PPort + portSpan
	lighthouseBeaconNodeMetricsPort = lighthouseBeaconNodeP2PPort + 2*portSpan

	validatorHTTPPort    = 6150
	validatorMetricsPort = validatorHTTPPort + portSpan

	jaegerTracingPort = 9150

	startupBuffer = 15 * time.Second
)

func logDir() string {
	wTime := func(p string) string {
		return path.Join(p, time.Now().Format("20060102/150405"))
	}
	path, ok := os.LookupEnv("E2E_LOG_PATH")
	if ok {
		return wTime(path)
	}
	path, _ = os.LookupEnv("TEST_UNDECLARED_OUTPUTS_DIR")
	return wTime(path)
}

// Init initializes the E2E config, properly handling test sharding.
func Init(t *testing.T, beaconNodeCount int) error {
	d := logDir()
	if d == "" {
		return errors.New("unable to determine log directory, no value for E2E_LOG_PATH or TEST_UNDECLARED_OUTPUTS_DIR")
	}
	logPath := path.Join(d, t.Name())
	if err := file.MkdirAll(logPath); err != nil {
		return err
	}
	testPath := bazel.TestTmpDir()
	testTotalShardsStr, ok := os.LookupEnv("TEST_TOTAL_SHARDS")
	if !ok {
		testTotalShardsStr = "1"
	}
	testTotalShards, err := strconv.Atoi(testTotalShardsStr)
	if err != nil {
		return err
	}
	testShardIndexStr, ok := os.LookupEnv("TEST_SHARD_INDEX")
	if !ok {
		testShardIndexStr = "0"
	}
	testShardIndex, err := strconv.Atoi(testShardIndexStr)
	if err != nil {
		return err
	}

	var existingRegistrations []int
	testPorts := &ports{}
	err = initializeStandardPorts(testTotalShards, testShardIndex, testPorts, &existingRegistrations)
	if err != nil {
		return err
	}

	genTime := time.Now().Add(startupBuffer)
	TestParams = &params{
		TestPath:               filepath.Join(testPath, fmt.Sprintf("shard-%d", testShardIndex)),
		LogPath:                logPath,
		TestShardIndex:         testShardIndex,
		BeaconNodeCount:        beaconNodeCount,
		Ports:                  testPorts,
		CLGenesisTime:          genTime,
		SilaExecutionGenesisTime:        genTime,
		NumberOfExecutionCreds: PregenesisExecCreds,
	}
	return nil
}

// InitMultiClient initializes the multiclient E2E config, properly handling test sharding.
func InitMultiClient(t *testing.T, beaconNodeCount int, lighthouseNodeCount int) error {
	testPath := bazel.TestTmpDir()
	logPath, ok := os.LookupEnv("TEST_UNDECLARED_OUTPUTS_DIR")
	if !ok {
		return errors.New("expected TEST_UNDECLARED_OUTPUTS_DIR to be defined")
	}
	logPath = path.Join(logPath, t.Name())
	if err := file.MkdirAll(logPath); err != nil {
		return err
	}
	testTotalShardsStr, ok := os.LookupEnv("TEST_TOTAL_SHARDS")
	if !ok {
		testTotalShardsStr = "1"
	}
	testTotalShards, err := strconv.Atoi(testTotalShardsStr)
	if err != nil {
		return err
	}
	testShardIndexStr, ok := os.LookupEnv("TEST_SHARD_INDEX")
	if !ok {
		testShardIndexStr = "0"
	}
	testShardIndex, err := strconv.Atoi(testShardIndexStr)
	if err != nil {
		return err
	}

	var existingRegistrations []int
	testPorts := &ports{}
	err = initializeStandardPorts(testTotalShards, testShardIndex, testPorts, &existingRegistrations)
	if err != nil {
		return err
	}
	err = initializeMulticlientPorts(testTotalShards, testShardIndex, testPorts, &existingRegistrations)
	if err != nil {
		return err
	}

	genTime := time.Now().Add(startupBuffer)
	TestParams = &params{
		TestPath:                  filepath.Join(testPath, fmt.Sprintf("shard-%d", testShardIndex)),
		LogPath:                   logPath,
		TestShardIndex:            testShardIndex,
		BeaconNodeCount:           beaconNodeCount,
		LighthouseBeaconNodeCount: lighthouseNodeCount,
		Ports:                     testPorts,
		CLGenesisTime:             genTime,
		SilaExecutionGenesisTime:           genTime,
		NumberOfExecutionCreds:    PregenesisExecCreds,
	}
	return nil
}

// port returns a safe port number based on the seed and shard data.
func port(seed, shardCount, shardIndex int, existingRegistrations *[]int) (int, error) {
	portToRegister := seed + portSpan/shardCount*shardIndex
	for _, p := range *existingRegistrations {
		if portToRegister >= p && portToRegister <= p+(portSpan/shardCount)-1 {
			return 0, fmt.Errorf("port %d overlaps with already registered port %d", seed, p)
		}
	}
	*existingRegistrations = append(*existingRegistrations, portToRegister)

	// Calculation example: 3 shards, seed 2000, port span 50.
	// Shard 0: 2000 + (50 / 3 * 0) = 2000 (we can safely use ports 2000-2015)
	// Shard 1: 2000 + (50 / 3 * 1) = 2016 (we can safely use ports 2016-2031)
	// Shard 2: 2000 + (50 / 3 * 2) = 2032 (we can safely use ports 2032-2047, and in reality 2032-2049)
	return portToRegister, nil
}

func initializeStandardPorts(shardCount, shardIndex int, ports *ports, existingRegistrations *[]int) error {
	bootnodePort, err := port(bootNodePort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	bootnodeMetricsPort, err := port(bootNodeMetricsPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	silaexecPort, err := port(silaexecPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	silaexecRPCPort, err := port(silaexecRPCPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	silaexecWSPort, err := port(silaexecWSPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	silaexecAuthPort, err := port(silaexecAuthRPCPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	silaexecProxyPort, err := port(silaexecProxyPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	beaconNodeRPCPort, err := port(silaBeaconNodeRPCPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	beaconNodeUDPPort, err := port(silaBeaconNodeUDPPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	beaconNodeQUICPort, err := port(silaBeaconNodeQUICPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	beaconNodeTCPPort, err := port(silaBeaconNodeTCPPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	beaconNodeHTTPPort, err := port(silaBeaconNodeHTTPPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	beaconNodeMetricsPort, err := port(silaBeaconNodeMetricsPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	beaconNodePprofPort, err := port(silaBeaconNodePprofPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	validatorHTTPPort, err := port(validatorHTTPPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	validatorMetricsPort, err := port(validatorMetricsPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	jaegerTracingPort, err := port(jaegerTracingPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	ports.BootNodePort = bootnodePort
	ports.BootNodeMetricsPort = bootnodeMetricsPort
	ports.SilaExecutionPort = silaexecPort
	ports.SilaExecutionRPCPort = silaexecRPCPort
	ports.SilaExecutionAuthRPCPort = silaexecAuthPort
	ports.SilaExecutionWSPort = silaexecWSPort
	ports.SilaExecutionProxyPort = silaexecProxyPort
	ports.SilaBeaconNodeRPCPort = beaconNodeRPCPort
	ports.SilaBeaconNodeUDPPort = beaconNodeUDPPort
	ports.SilaBeaconNodeQUICPort = beaconNodeQUICPort
	ports.SilaBeaconNodeTCPPort = beaconNodeTCPPort
	ports.SilaBeaconNodeHTTPPort = beaconNodeHTTPPort
	ports.SilaBeaconNodeMetricsPort = beaconNodeMetricsPort
	ports.SilaBeaconNodePprofPort = beaconNodePprofPort
	ports.ValidatorMetricsPort = validatorMetricsPort
	ports.ValidatorHTTPPort = validatorHTTPPort
	ports.JaegerTracingPort = jaegerTracingPort
	return nil
}

func initializeMulticlientPorts(shardCount, shardIndex int, ports *ports, existingRegistrations *[]int) error {
	lighthouseBeaconNodeP2PPort, err := port(lighthouseBeaconNodeP2PPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	lighthouseBeaconNodeHTTPPort, err := port(lighthouseBeaconNodeHTTPPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	lighthouseBeaconNodeMetricsPort, err := port(lighthouseBeaconNodeMetricsPort, shardCount, shardIndex, existingRegistrations)
	if err != nil {
		return err
	}
	ports.LighthouseBeaconNodeP2PPort = lighthouseBeaconNodeP2PPort
	ports.LighthouseBeaconNodeHTTPPort = lighthouseBeaconNodeHTTPPort
	ports.LighthouseBeaconNodeMetricsPort = lighthouseBeaconNodeMetricsPort
	return nil
}
