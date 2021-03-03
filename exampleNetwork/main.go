package main

import (
	"blockchain_example/blockchain"
	"blockchain_example/network"
	"blockchain_example/node"
	"crypto/sha256"
	"fmt"
	"sync"
)

const RequiredLeadingZeros = 3

var hashFactory blockchain.HashFactory

func printChainStatus(chains []blockchain.Blockchain) {
	fmt.Printf("Currently number of chains %d\n", len(chains))
	for _, chain := range chains {
		printChain(chain)
	}
	fmt.Println("Most valid chain")
	printChain(blockchain.GetMostValidBlockChain(chains, RequiredLeadingZeros, hashFactory))
}

func printChain(chain blockchain.Blockchain) {
	fmt.Printf("%d ", len(chain))
	for _, block := range chain {
		fmt.Printf("(%d, %v), ", block.Index, block.Data)
	}
	fmt.Println("")
}

func main() {
	hashFactory = sha256.New
	fmt.Println("Hello mining")
	network := network.Network{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	network.SubscribeForNewBlock(func(_ blockchain.Block) {
		wg.Done()
	})

	node.NewNode("First", &network, RequiredLeadingZeros, hashFactory)

	// let the first node get the first block.
	network.PostNewData("a")
	wg.Wait()

	printChainStatus(network.GetBlochains())

	node.NewNode("Second", &network, RequiredLeadingZeros, hashFactory)

	simulatedData := []string{"b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for _, data := range simulatedData {
		fmt.Printf("\nSimulator sending data %v\n", data)
		wg.Add(1)
		network.PostNewData(data)
		wg.Wait()

		printChainStatus(network.GetBlochains())
	}
}
