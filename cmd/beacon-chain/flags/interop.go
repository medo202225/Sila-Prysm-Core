package flags

import (
	"github.com/urfave/cli/v2"
)

var (
	// InteropMockSilaExecutionDataVotesFlag enables mocking the silaexec proof-of-work chain data put into blocks by proposers.
	InteropMockSilaExecutionDataVotesFlag = &cli.BoolFlag{
		Name:  "interop-silaExecutionData-votes",
		Usage: "Enable mocking of silaexec data votes for proposers to package into blocks",
	}
)
