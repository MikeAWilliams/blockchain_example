package network_test

import (
	"blockchain_example/blockchain"
	"blockchain_example/network"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewData(t *testing.T) {
	testObject := &network.Network{}
	f1Called := 0
	var f1Arg string
	testObject.SubscribeForNewData(func(data string) {
		f1Called++
		f1Arg = data
	})
	f2Called := 0
	var f2Arg string
	testObject.SubscribeForNewData(func(data string) {
		f2Called++
		f2Arg = data
	})

	testObject.PostNewData("TACO")

	require.Equal(t, 1, f1Called)
	require.Equal(t, 1, f2Called)
	require.Equal(t, "TACO", f1Arg)
	require.Equal(t, "TACO", f2Arg)
}

func Test_NewBlock(t *testing.T) {
	testObject := &network.Network{}
	f1Called := 0
	var f1Arg blockchain.Block
	testObject.SubscribeForNewBlock(func(arg blockchain.Block) {
		f1Called++
		f1Arg = arg
	})
	f2Called := 0
	var f2Arg blockchain.Block
	testObject.SubscribeForNewBlock(func(arg blockchain.Block) {
		f2Called++
		f2Arg = arg
	})

	expectedBlock := blockchain.Block{Index: 12, Data: "Taco"}
	testObject.PostNewBlock(expectedBlock)

	require.Equal(t, 1, f1Called)
	require.Equal(t, 1, f2Called)
	require.Equal(t, expectedBlock, f1Arg)
	require.Equal(t, expectedBlock, f2Arg)
}

func Test_GetBlockchains(t *testing.T) {
	testObject := &network.Network{}
	testObject.RegisterAsBlockchainProvider(func() blockchain.Blockchain {
		return blockchain.Blockchain{blockchain.Block{}, blockchain.Block{}, blockchain.Block{}}
	})
	testObject.RegisterAsBlockchainProvider(func() blockchain.Blockchain {
		return blockchain.Blockchain{blockchain.Block{}, blockchain.Block{}}
	})
	testObject.RegisterAsBlockchainProvider(func() blockchain.Blockchain {
		return blockchain.Blockchain{blockchain.Block{}}
	})

	result := testObject.GetBlochains()

	require.Equal(t, 3, len(result))
	require.True(t, 3 == len(result[0]) || 3 == len(result[1]) || 3 == len(result[2]))
	require.True(t, 2 == len(result[0]) || 2 == len(result[1]) || 2 == len(result[2]))
	require.True(t, 1 == len(result[0]) || 1 == len(result[1]) || 1 == len(result[2]))
}
