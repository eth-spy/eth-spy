package main

import (
	"net"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/p2p"
	"log"
)

// Struct for handling all the necessary configuration options for our ETHSpy instance
// TODO: Define what such a struct should look like
type EthSpy struct {

}

// Start : Starts an instance of the ETHSpy client
// TODO : Implement Start
func (es *EthSpy) Start() {
	// TODO: Connect to a specific node or run peer discovery protocol

	// TODO: If connected to a specific node, request nodes in its RLPx DHT 
	// and recursively get those nodes's DHTs as well

	// TODO: Store discovered nodes 

	// TODO: Connect to discovered nodes using the RLPx handshake
	
	// TODO: Ensure that the minimum is done to keep connections alive
	// For our purposes, we need to only consider transaction exchanges (later we can also consider block propagation)
	// Things that may get us disconnected:
	//     - Sending protocol messages that are larger than the limits
	//     - Propagating invalid transactions
	//     - Propagating empty transactions

	// TODO: Upon connecting to these nodes, listen for any incoming transactions
	// When listening for transactions, record the following information:
	//     - IP address of the node
	//     - Timestamp of when the transaction was received
	//     - Transaction Hash of the transaction

}

// discoveryHandler : Handles discovering nodes on the network
func (es *EthSpy) discoveryHandler() {

}

// transactionHandler : Handlers waiting and relaying transactions on the network
func (es *EthSpy) transactionHandler() {

}

// Stop : Disconnects the ETHSpy client from all of its connected peers
func (es *EthSpy) Stop() {
	// TODO : Stop running discovery
	// TODO : Stop running transaction listening and relaying
	// TODO : Clean up any left of threads or processes

}