package blockchain_test

import (
	"blockchain_example/blockchain"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BlockIsValid_Happy(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 5, PreviousHash: block1.Hash, Hash: []byte{'0', '0', 'q'}}

	require.True(t, blockchain.IsValid(block1, block2, 2))
}

func Test_BlockIsValid_Stupid(t *testing.T) {
	block1 := blockchain.Block{}
	block2 := blockchain.Block{}

	require.False(t, blockchain.IsValid(block1, block2, 0))
}
