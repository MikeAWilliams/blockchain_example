package blockchain_test

import (
	"blockchain_example/blockchain"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BlockIsValid_Happy(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 5, PreviousHash: block1.Hash, Hash: []byte{'0', '0', 'q'}}

	hasher := fixedHash{hashResult: block2.Hash}
	require.True(t, blockchain.IsValid(block1, block2, 2, hasher))
}

func Test_BlockIsValid_BadIndex(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 6, PreviousHash: block1.Hash, Hash: []byte{'0', '0', 'q'}}

	hasher := fixedHash{hashResult: block2.Hash}
	require.False(t, blockchain.IsValid(block1, block2, 0, hasher))
}

func Test_BlockIsValid_BadDoesNotMatchPrevious(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 5, PreviousHash: []byte{'0', '0', 'q'}}

	hasher := fixedHash{hashResult: block2.Hash}
	require.False(t, blockchain.IsValid(block1, block2, 0, hasher))
}

func Test_BlockIsValid_HashInvalidLeadingZero(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 5, PreviousHash: block1.Hash}

	hasher := fixedHash{hashResult: block2.Hash}
	require.False(t, blockchain.IsValid(block1, block2, 3, hasher))
}

func Test_BlockIsValid_HashIncorrect(t *testing.T) {
	block1 := blockchain.Block{Index: 4, Hash: []byte{'0', '0', 'n'}}
	block2 := blockchain.Block{Index: 5, PreviousHash: block1.Hash, Hash: []byte{'0', '0', 'a'}}

	hasher := fixedHash{hashResult: []byte{'0', '0', 'q'}}
	require.False(t, blockchain.IsValid(block1, block2, 2, hasher))
}

func Test_FirstBlockIsValid_Happy(t *testing.T) {
	block := blockchain.Block{Index: 0, PreviousHash: nil, Hash: []byte{'0', '0', 'q'}}
	hasher := fixedHash{hashResult: block.Hash}

	require.True(t, blockchain.FirstBlockIsValid(block, 2, hasher))
}

func Test_FirstBlockIsValid_BadIndex(t *testing.T) {
	block := blockchain.Block{Index: 1, PreviousHash: nil, Hash: []byte{'0', '0', 'q'}}
	hasher := fixedHash{hashResult: block.Hash}

	require.False(t, blockchain.FirstBlockIsValid(block, 2, hasher))
}

func Test_FirstBlockIsValid_BadPreviousHash(t *testing.T) {
	block := blockchain.Block{Index: 0, PreviousHash: []byte{}, Hash: []byte{'0', '0', 'q'}}
	hasher := fixedHash{hashResult: block.Hash}

	require.False(t, blockchain.FirstBlockIsValid(block, 2, hasher))
}

func Test_FirstBlockIsValid_HashNotHardEnough(t *testing.T) {
	block := blockchain.Block{Index: 0, PreviousHash: nil, Hash: []byte{'0', '0', 'q'}}
	hasher := fixedHash{hashResult: block.Hash}

	require.False(t, blockchain.FirstBlockIsValid(block, 3, hasher))
}

func Test_FirstBlockIsValid_HashIncorrect(t *testing.T) {
	block := blockchain.Block{Index: 0, PreviousHash: nil, Hash: []byte{'0', '0', 'c'}}

	hasher := fixedHash{hashResult: []byte{'0', '0', 'q'}}
	require.False(t, blockchain.FirstBlockIsValid(block, 2, hasher))
}
