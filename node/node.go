package node

import (
	"blockchain_example/blockchain"
	"blockchain_example/network"
	"fmt"
	"strings"
)

type Node struct {
	network              *network.Network
	name                 string
	hashFactory          blockchain.HashFactory
	requiredLeadingZeros int
	chain                blockchain.Blockchain
}

func NewNode(name string, network *network.Network, requiredZeros int, hashFactory blockchain.HashFactory) Node {
	//result := Node{network: network, name: name, requiredLeadingZeros: requiredZeros, hashFactory: hashFactory}
	// since I am using the sleep mining there is no garontee of the leading zeros so setting it to 0 will let validate work
	result := Node{network: network, name: name, requiredLeadingZeros: 0, hashFactory: hashFactory}
	result.chain = blockchain.Blockchain{}
	result.RegisterToNetwork()
	return result
}

func (n *Node) NewData(data string) {
	go n.mine(n.name + "-" + data)
}

func (n *Node) mine(data string) {
	fmt.Printf("%v mining data %v\n", n.name, data)
	var previousHash []byte
	if 0 == len(n.chain) {
		previousHash = nil
	} else {
		previousHash = n.chain[len(n.chain)-1].Hash
	}

	newBlock, _ := blockchain.MineBlockWithSleep(previousHash, int64(len(n.chain)-1), data, 5, 25, n.hashFactory)
	//newBlock, _ := blockchain.MineBlock(previousHash, int64(len(n.chain)-1), data, n.requiredLeadingZeros, n.hashFactory)
	if int64(len(n.chain)) > newBlock.Index {
		return
	}
	n.chain = append(n.chain, newBlock)
	n.network.PostNewBlock(newBlock)
}

func (n *Node) NewBlock(block blockchain.Block) {
	if strings.Contains(block.Data, n.name) {
		return
	}
	if blockchain.IsValid(n.chain[len(n.chain)-1], block, n.requiredLeadingZeros, n.hashFactory()) {
		n.chain = append(n.chain, block)
	} else {
		fmt.Printf("--------- %v recieved invalid block with data %v ------- \n", n.name, block.Data)
	}
}

func (n *Node) GetBlockchain() blockchain.NamedChain {
	return blockchain.NamedChain{Name: n.name, Chain: n.chain}
}

func (n *Node) InitilizeBlockChain() {
	allChains := n.network.GetBlochains()
	n.chain = blockchain.GetMostValidBlockChain(allChains, n.requiredLeadingZeros, n.hashFactory).Chain
}

func (n *Node) RegisterToNetwork() {
	n.network.SubscribeForNewData(n.NewData)
	n.network.SubscribeForNewBlock(n.NewBlock)
	n.InitilizeBlockChain()
	n.network.RegisterAsBlockchainProvider(n.GetBlockchain)
}
