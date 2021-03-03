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
	testObject.RegisterAsBlockchainProvider(func() blockchain.NamedChain {
		return blockchain.NamedChain{Name: "1", Chain: blockchain.Blockchain{blockchain.Block{}, blockchain.Block{}, blockchain.Block{}}}
	})
	testObject.RegisterAsBlockchainProvider(func() blockchain.NamedChain {
		return blockchain.NamedChain{Name: "2", Chain: blockchain.Blockchain{blockchain.Block{}, blockchain.Block{}}}
	})
	testObject.RegisterAsBlockchainProvider(func() blockchain.NamedChain {
		return blockchain.NamedChain{Name: "3", Chain: blockchain.Blockchain{blockchain.Block{}}}
	})

	result := testObject.GetBlochains()

	require.Equal(t, 3, len(result))
	oneFound := false
	twoFound := false
	threeFound := false
	for _, nameChain := range result {
		if "1" == nameChain.Name {
			require.Equal(t, 3, len(nameChain.Chain))
			oneFound = true
		}
		if "2" == nameChain.Name {
			require.Equal(t, 2, len(nameChain.Chain))
			twoFound = true
		}
		if "3" == nameChain.Name {
			require.Equal(t, 1, len(nameChain.Chain))
			threeFound = true
		}
	}
	require.True(t, oneFound)
	require.True(t, twoFound)
	require.True(t, threeFound)
}
