package node

import (
	"blockchain_example/blockchain"
	"blockchain_example/network"
	"fmt"
	"strings"
	"sync"
)

type Node struct {
	network              *network.Network
	name                 string
	hashFactory          blockchain.HashFactory
	requiredLeadingZeros int
	chain                blockchain.Blockchain
	chain2               blockchain.Blockchain
	mutex                sync.Mutex
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
	n.mutex.Lock()
	if int64(len(n.chain)) > newBlock.Index {
		n.mutex.Unlock()
		return
	}
	n.chain = append(n.chain, newBlock)
	n.mutex.Unlock()
	n.network.PostNewBlock(newBlock)
}

func (n *Node) tryToConsolateChains() {
	if len(n.chain) > len(n.chain2) {
		n.chain2 = nil
		return
	}
	if len(n.chain2) > len(n.chain) {
		n.chain = n.chain2
		n.chain2 = nil
	}
}

func (n *Node) createNewChainIfNeeded(block blockchain.Block) bool {
	if nil != n.chain2 {
		return false
	}
	if block.Index == int64(len(n.chain)-1) {
		if blockchain.IsValid(n.chain[len(n.chain)-2], block, n.requiredLeadingZeros, n.hashFactory()) {
			n.chain2 = append(n.chain2, n.chain...)
			n.chain2[len(n.chain2)-1] = block
			fmt.Printf("--------- %v created new chain with data %v------- \n", n.name, block.Data)
			return true
		}
		fmt.Printf("--------- %v recieved invalid block with data %v while making new chain------- \n", n.name, block.Data)
		return true
	}
	return false
}

func (n *Node) NewBlock(block blockchain.Block) {
	if strings.Contains(block.Data, n.name) {
		return
	}
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if n.createNewChainIfNeeded(block) {
		return
	}
	if blockchain.IsValid(n.chain[len(n.chain)-1], block, n.requiredLeadingZeros, n.hashFactory()) {
		n.chain = append(n.chain, block)
	} else if nil != n.chain2 && blockchain.IsValid(n.chain2[len(n.chain2)-1], block, n.requiredLeadingZeros, n.hashFactory()) {
		n.chain2 = append(n.chain2, block)
	} else {
		fmt.Printf("--------- %v recieved invalid block with data %v ------- \n", n.name, block.Data)
	}
	n.tryToConsolateChains()
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
