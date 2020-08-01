package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	startCmd := flag.NewFlagSet("start", flag.ExitOnError)
	startOnlyTxs := startCmd.Bool("only-txs", true, "Only start logging transactions")
	startOnlyBlocks := startCmd.Bool("only-blocks", false, "Only start logging blocks")
	startBothTxsAndBlocks := startCmd.Bool("txs-and-blocks", false, "Start logging both transactions and blocks")
	startMaxPeers := startCmd.Int("max-peers", 50, "Max number of peers that can be connected")
	startMaxPendingPeers := startCmd.Int("max-pending-peers", 50, "Max number of peers that can be pending in the handshake phase")
	startListeningAddr := startCmd.String("listening-addr", "", "Address at which eth-spy should listen at")
	// disconnectPtr := flag.Bool("disconnect", false, "Disconnects the ETHspy client from its victims")
	disconnectCmd := flag.NewFlagSet("disconnect", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Errorf("Expected either to 'start' or to 'disconnect' ETHspy")
		os.Exit(1)
	}

	ethspy := &EthSpy{}

	switch os.Args[1] {
	case "start":
		startCmd.Parse(os.Args[2:])
		fmt.Println("Starting ETHspy ...")
		// TODO: Call our ETHspy client with the provided command line arguments
		fmt.Println("only-tx, ", *startOnlyTxs)
		fmt.Println("only-blocks, ", *startOnlyBlocks)
		fmt.Println("Txs and Blocks, ", *startBothTxsAndBlocks)
		
		ethspy, err := NewEthSpy(*startMaxPeers, *startMaxPendingPeers, *startListeningAddr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		err = ethspy.Start()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "disconnect":
		disconnectCmd.Parse(os.Args[2:])
		fmt.Println("Disconnecting ETHspy ...")
		ethspy.Stop()
	default:
		fmt.Errorf("Expected either to 'start' or to 'disconnect' ETHspy")
		os.Exit(1)
	}
}