package main

import (
	"blockchain_example/blockchain"
	"blockchain_example/network"
	"blockchain_example/node"
	"crypto/sha256"
	"fmt"
	"strconv"
	"sync"
)

const RequiredLeadingZeros = 3

func main() {
	fmt.Println("Hello mining")
	network := network.Network{}
	hashFactory := sha256.New

	wg := sync.WaitGroup{}
	wg.Add(1)
	network.SubscribeForNewBlock(func(_ blockchain.Block) {
		wg.Done()
	})

	node.NewNode("First", &network, RequiredLeadingZeros, hashFactory)

	// let the first node get the first block.
	network.PostNewData("a")
	wg.Wait()

	node.NewNode("Second", &network, RequiredLeadingZeros, hashFactory)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		network.PostNewData(strconv.Itoa(i))
		wg.Wait()
	}
}
