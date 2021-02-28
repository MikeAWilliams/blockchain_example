package blockchain_test

import (
	"blockchain_example/blockchain"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BlockIsValid_stupid(t *testing.T) {
	block1 := blockchain.Block{}
	block2 := blockchain.Block{}

	require.False(t, blockchain.IsValid(block1, block2))
}
