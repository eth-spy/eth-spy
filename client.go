package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"time"
)


// Struct for handling all the necessary configuration options for our ETHSpy instance
// TODO: Define what such a struct should look like
type EthSpy struct {
	p2pserver	*p2p.Server
}

func newkey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic("couldn't generate key: " + err.Error())
	}
	fmt.Println("ID", enode.PubkeyToIDV4(&key.PublicKey))
	return key
}

func (es *EthSpy) AddPeer(nodeStr string) {
	node, err := enode.ParseV4(nodeStr)
	if err != nil {
		fmt.Println(err)
	}
	es.p2pserver.AddPeer(node)
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

	if err := es.p2pserver.Start(); err != nil {
		fmt.Println(err)
	}
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

// NewEthSpy : Creates a new instance of an EthSpy struct
func NewEthSpy() (*EthSpy, error) {
	// start the server
	//srv := startServer()

	config := p2p.Config{
		MaxPeers:    10,
		PrivateKey:  newkey(),
	}
	server := &p2p.Server{
		Config:       config,
	}
	for v := 65; v < 66; v++ {
		server.Protocols = append(server.Protocols, p2p.Protocol{Name: "eth", Version:uint(v)})
	}

	ethspy := EthSpy{
		p2pserver:	server,
	}
	return &ethspy, nil
}

func main() {
	ethspy, _ := NewEthSpy()
	ethspy.Start()
	fmt.Println("nodeInfo", ethspy.p2pserver.NodeInfo())
	defer ethspy.p2pserver.Stop()
	enodeRawUrl := "enode://5773aafffd6b5c8c383c5000a97cf14b82c550b2a1aab26a9d38d74f0c5309360634a626d9e8a76dccee652a2e9544cbed92e3ad921d18a6cec15930a9b664c4@127.0.0.1:30304"
	ethspy.AddPeer(enodeRawUrl)
	for i := 0; i < 60; i++ {
		time.Sleep(time.Second)
		fmt.Println(ethspy.p2pserver.PeerCount())
	}
}