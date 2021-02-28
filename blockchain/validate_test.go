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

func Test_BlockIsValid_BadIndex(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 6, PreviousHash: block1.Hash, Hash: []byte{'0', '0', 'q'}}

	require.False(t, blockchain.IsValid(block1, block2, 0))
}

func Test_BlockIsValid_BadDoesNotMatchPrevious(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 5, PreviousHash: []byte{'0', '0', 'q'}}

	require.False(t, blockchain.IsValid(block1, block2, 0))
}

func Test_BlockIsValid_HashInvalid(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 5, PreviousHash: block1.Hash}

	require.False(t, blockchain.IsValid(block1, block2, 3))
}
