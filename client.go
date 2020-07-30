package main

import (
	"crypto/ecdsa"
	"fmt"
	gethCrypto "github.com/ethereum/go-ethereum/crypto"
	params "github.com/ethereum/go-ethereum/params"
	p2p "github.com/ethereum/go-ethereum/p2p"
	enode "github.com/ethereum/go-ethereum/p2p/enode"
	
)

// Struct for handling all the necessary configuration options for our ETHSpy instance
// TODO: Define what such a struct should look like
type EthSpy struct {
	privateKey *ecdsa.PrivateKey
	server *p2p.Server

}

// Start : Starts an instance of the ETHSpy client
// TODO : Implement Start
func (es *EthSpy) Start() error {
	// TODO: Connect to a specific node or run peer discovery protocol

	// TODO: If connected to a specific node, request nodes in its RLPx DHT 
	// and recursively get those nodes's DHTs as well

	fmt.Println("Starting ETHspy Server...")
	es.server.Start()
	return nil

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
	es.server.Stop()
}

// NewEthSpy : Creates a new instance of an EthSpy struct
func NewEthSpy(maxPeers int, maxPendingPeers int) (*EthSpy, error) {

	ethspy := &EthSpy{}
	privateKey, err := gethCrypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	// Parse bootnodes
	nodes := make([]*enode.Node, len(params.MainnetBootnodes))
	for i, url := range params.MainnetBootnodes {
		nodes[i] = enode.MustParse(url)
	}

	config := p2p.Config{
		PrivateKey: privateKey,
		MaxPeers: maxPeers,
		MaxPendingPeers: maxPendingPeers,
		DialRatio: 3,
		NoDiscovery: false,
		Name: "eth-spy",
		BootstrapNodes: nodes,
		NAT: nil,
		NoDial: false,
		EnableMsgEvents: true,
	}

	server := &p2p.Server{Config: config}

	ethspy.privateKey = privateKey
	ethspy.server = server

	return ethspy, nil
}