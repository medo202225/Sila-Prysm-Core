package main

import (
	"os"

	"github.com/sila-chain/Sila-Consensus-Core/v7/cmd/silactl/checkpointsync"
	"github.com/sila-chain/Sila-Consensus-Core/v7/cmd/silactl/db"
	"github.com/sila-chain/Sila-Consensus-Core/v7/cmd/silactl/p2p"
	"github.com/sila-chain/Sila-Consensus-Core/v7/cmd/silactl/testnet"
	"github.com/sila-chain/Sila-Consensus-Core/v7/cmd/silactl/validator"
	"github.com/sila-chain/Sila-Consensus-Core/v7/cmd/silactl/weaksubjectivity"
	"github.com/urfave/cli/v2"
)

var silactlCommands []*cli.Command

func main() {
	app := &cli.App{
		Commands: silactlCommands,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	silactlCommands = append(silactlCommands, checkpointsync.Commands...)
	silactlCommands = append(silactlCommands, db.Commands...)
	silactlCommands = append(silactlCommands, p2p.Commands...)
	silactlCommands = append(silactlCommands, testnet.Commands...)
	silactlCommands = append(silactlCommands, weaksubjectivity.Commands...)
	silactlCommands = append(silactlCommands, validator.Commands...)
}
