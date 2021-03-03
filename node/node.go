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
	go n.mine(data)
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
	n.chain = append(n.chain, newBlock)
	n.network.PostNewBlock(newBlock)
}

func (n *Node) NewBlock(block blockchain.Block) {
	go func() {
		fmt.Printf("%v recieved block with data %v\n", n.name, block.Data)
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

func (n *Node) getValidBlockchains(chains []blockchain.Blockchain) []blockchain.Blockchain {
	result := []blockchain.Blockchain{}

	for _, chain := range chains {
		if blockchain.ChainIsValid(chain, n.requiredLeadingZeros, n.hashFactory) {
			result = append(result, chain)
		}
	}

	return result
}

func getMostValidBlockchain(validChains []blockchain.Blockchain) blockchain.Blockchain {
	var longest blockchain.Blockchain
	for _, chain := range validChains {
		if len(chain) > len(longest) {
			longest = chain
		}
	}
	return longest
}

func (n *Node) InitilizeBlockChain() {
	allChains := n.network.GetBlochains()
	validChains := n.getValidBlockchains(allChains)
	n.chain = getMostValidBlockchain(validChains)
}

func (n *Node) RegisterToNetwork() {
	n.network.SubscribeForNewData(n.NewData)
	n.network.SubscribeForNewBlock(n.NewBlock)
	n.InitilizeBlockChain()
	n.network.RegisterAsBlockchainProvider(n.GetBlockchain)
}