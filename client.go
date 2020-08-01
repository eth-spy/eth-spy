package main

import (
	"crypto/ecdsa"
	"fmt"
	"net"
	gethCrypto "github.com/ethereum/go-ethereum/crypto"
	node "github.com/ethereum/go-ethereum/node"
	params "github.com/ethereum/go-ethereum/params"
	//eth "github.com/ethereum/go-ethereum/eth"
	p2p "github.com/ethereum/go-ethereum/p2p"
	//eth "github.com/ethereum/go-ethereum/eth"
	enode "github.com/ethereum/go-ethereum/p2p/enode"
	discover "github.com/ethereum/go-ethereum/p2p/discover"
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

	utils.StartNode(es.node)

	fmt.Println("Starting ETHspy Server...")
	server := es.node.Server()
	//fmt.Println("server listen addr: %v", server.Config.ListenAddr)
	
	addr, err := net.ResolveUDPAddr("udp", server.ListenAddr)
	if err != nil {
		return fmt.Errorf("Cannot resolved UDP address: %v", err)
	}
	
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("Cannot listen at given address: %v", err)
	}
	
	var unhandled chan discover.ReadPacket
	discoverConfig := discover.Config{
		PrivateKey : server.PrivateKey,
		NetRestrict : server.NetRestrict,
		Bootnodes : server.BootstrapNodes,
		Unhandled : unhandled,
		Log : server.Config.Logger,
	}
	ntab, err := discover.ListenUDP(conn, server.LocalNode(), discoverConfig)
	if err != nil {
		return fmt.Errorf("Cannot listen for discovery connections: %v", err)
	}
	
	nodeIterator := ntab.RandomNodes()
	for nodeIterator.Next() {
		fmt.Println("node: %v", nodeIterator.Node())
	}

	return nil
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
func NewEthSpy(maxPeers int, maxPendingPeers int, listeningAddr string) (*EthSpy, error) {

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
	nodeConfig.P2P.Protocols = []p2p.Protocol{pingProtocol, pongProtocol}
	nodeConfig.P2P.EnableMsgEvents = true
	nodeConfig.P2P.ListenAddr = listeningAddr
	node, err := node.New(&nodeConfig)
	if err != nil {
		return nil, err
	}

	//utils.RegisterEthService(node, &eth.DefaultConfig)

	ethspy.privateKey = privateKey
	//ethspy.server = server
	ethspy.node = node

	return ethspy, nil
}