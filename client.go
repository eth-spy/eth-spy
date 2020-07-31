package main

import (
	"crypto/ecdsa"
	"fmt"
	gethCrypto "github.com/ethereum/go-ethereum/crypto"
	node "github.com/ethereum/go-ethereum/node"
	params "github.com/ethereum/go-ethereum/params"
	//eth "github.com/ethereum/go-ethereum/eth"
	//p2p "github.com/ethereum/go-ethereum/p2p"
	eth "github.com/ethereum/go-ethereum/eth"
	enode "github.com/ethereum/go-ethereum/p2p/enode"
	utils "github.com/ethereum/go-ethereum/cmd/utils"
	
)

// Struct for handling all the necessary configuration options for our ETHSpy instance
// TODO: Define what such a struct should look like
type EthSpy struct {
	privateKey *ecdsa.PrivateKey
	// server *p2p.Server
	node *node.Node

}

// Start : Starts an instance of the ETHSpy client
// TODO : Implement Start
func (es *EthSpy) Start() error {
	// TODO: Connect to a specific node or run peer discovery protocol

	// TODO: If connected to a specific node, request nodes in its RLPx DHT 
	// and recursively get those nodes's DHTs as well

	fmt.Println("Starting ETHspy Server...")
	//err := es.server.Start()
	

	utils.StartNode(es.node)
	//if err != nil {
	//	return err
	//}

	fmt.Println("NodeInfo: ", es.node.Server().NodeInfo())
	fmt.Println("NodeDB: ", es.node.Config().NodeDB())
	es.node.Wait()

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
	//es.server.Stop()
	es.node.Close()
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

	//p2pConfig := p2p.Config{
	//	PrivateKey: privateKey,
	//	MaxPeers: maxPeers,
	//	MaxPendingPeers: maxPendingPeers,
	//	Name: "eth-spy",
	//	BootstrapNodes: nodes,
	//}

	//nodeConfig := &node.Config{
	//	Name: "eth-spy",
	//	Version: "0.0.1",
	//	P2P: p2pConfig,
	//}

	//server := &p2p.Server{Config: p2pConfig}
	nodeConfig := node.DefaultConfig
	nodeConfig.P2P.StaticNodes = nodes
	nodeConfig.P2P.TrustedNodes = nodes
	nodeConfig.P2P.BootstrapNodes = nodes
	node, err := node.New(&nodeConfig)
	if err != nil {
		return nil, err
	}

	utils.RegisterEthService(node, &eth.DefaultConfig)

	ethspy.privateKey = privateKey
	//ethspy.server = server
	ethspy.node = node

	return ethspy, nil
}