package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	startCmd := flag.NewFlagSet("start", flag.ExitOnError)
	startOnlyTxs := startCmd.Bool("only-txs", true, "Only start logging transactions")
	startOnlyBlocks := startCmd.Bool("only-txs", false, "Only start logging blocks")
	startBothTxsAndBlocks := startCmd.Bool("txs-and-blocks", false, "Start logging both transactions and blocks")
	startEthClient := startCmd.String("eth-client", "", "Specific Ethereum client to connect to")
	disconnectPtr := flag.Bool("disconnect", false, "Disconnects the ETHspy client from its victims")

	if len(os.Args) < 2 {
		fmt.Println("Expected either to 'start' or to 'disconnect' ETHspy")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "start":
		startCmd.Parse(os.Args[2:])
		fmt.Println("Starting ETHspy ...")
		// TODO: Call our ETHspy client with the provided command line arguments
	case "disconnect":
		flag.Parse()
		fmt.Println("Disconnecting ETHspy ...")
		// TODO: Call to the disconnect functionality in our ETHspy client
	default:
		fmt.Println("Expected either to 'start' or to 'disconnect' ETHspy")
		os.Exit(1)
	}
}