package node

import (
	"blockchain_example/blockchain"
	"blockchain_example/network"
	"fmt"
)

type Node struct {
	network              *network.Network
	name                 string
	hashFactory          blockchain.HashFactory
	requiredLeadingZeros int
	chain                blockchain.Blockchain
}

func NewNode(name string, network *network.Network, requiredZeros int, hashFactory blockchain.HashFactory) Node {
	result := Node{network: network, name: name, requiredLeadingZeros: requiredZeros, hashFactory: hashFactory}
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

	newBlock, _ := blockchain.MineBlock(previousHash, int64(len(n.chain)-1), data, n.requiredLeadingZeros, n.hashFactory)
	if 0 == len(n.chain) || blockchain.IsValid(n.chain[len(n.chain)-1], newBlock, n.requiredLeadingZeros, n.hashFactory()) {
		n.chain = append(n.chain, newBlock)
		n.network.PostNewBlock(newBlock)
	}
}

func (n *Node) NewBlock(block blockchain.Block) {
	go func() {
		fmt.Printf("%v recieved block with data %v\n", n.name, block.Data)
		if block.Index < int64(len(n.chain)) {
			fmt.Printf("%v decided block with data %v is from the past\n", n.name, block.Data)
			return
		}
		if blockchain.IsValid(n.chain[len(n.chain)-1], block, n.requiredLeadingZeros, n.hashFactory()) {
			fmt.Printf("%v decided block with data %v is valid\n", n.name, block.Data)
			n.chain = append(n.chain, block)
		} else {
			fmt.Printf("%v decided block with data %v is not valid\n", n.name, block.Data)
		}

	}()
}

func (n *Node) GetBlockchain() blockchain.Blockchain {
	return n.chain
}

func (n *Node) InitilizeBlockChain() {
	allChains := n.network.GetBlochains()
	n.chain = blockchain.GetMostValidBlockChain(allChains, n.requiredLeadingZeros, n.hashFactory)
}

func (n *Node) RegisterToNetwork() {
	n.network.SubscribeForNewData(n.NewData)
	n.network.SubscribeForNewBlock(n.NewBlock)
	n.InitilizeBlockChain()
	n.network.RegisterAsBlockchainProvider(n.GetBlockchain)
}
