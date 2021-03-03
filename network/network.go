package network

type NewDataCallback func(string)

type Network struct {
	newDataSubscribers []NewDataCallback
}

func (n *Network) PostNewData(data string) {
	for _, callback := range n.newDataSubscribers {
		callback(data)
	}
}

func (n *Network) SubscribeForNewData(callback NewDataCallback) {
	n.newDataSubscribers = append(n.newDataSubscribers, callback)
}
