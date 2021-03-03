package network

import (
	"blockchain_example/blockchain"
	"math/rand"
)

type NewDataCallback func(string)
type NewBlockCallback func(blockchain.Block)

type Network struct {
	newDataSubscribers  []NewDataCallback
	newBlockSubscribers []NewBlockCallback
}

func (n *Network) PostNewData(data string) {
	rand.Shuffle(len(n.newDataSubscribers), func(i, j int) {
		n.newDataSubscribers[i], n.newDataSubscribers[j] = n.newDataSubscribers[j], n.newDataSubscribers[i]
	})
	for _, callback := range n.newDataSubscribers {
		callback(data)
	}
}

func (n *Network) SubscribeForNewData(callback NewDataCallback) {
	n.newDataSubscribers = append(n.newDataSubscribers, callback)
}

func (n *Network) PostNewBlock(data blockchain.Block) {
	rand.Shuffle(len(n.newBlockSubscribers), func(i, j int) {
		n.newBlockSubscribers[i], n.newBlockSubscribers[j] = n.newBlockSubscribers[j], n.newBlockSubscribers[i]
	})
	for _, callback := range n.newBlockSubscribers {
		callback(data)
	}
}

func (n *Network) SubscribeForNewBlock(callback NewBlockCallback) {
	n.newBlockSubscribers = append(n.newBlockSubscribers, callback)
}
