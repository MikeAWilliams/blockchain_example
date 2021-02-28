package blockchain_test

import (
	"blockchain_example/blockchain"
	"crypto/sha256"
	"fmt"
	"hash"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HashIsHardEnough(t *testing.T) {
	badData := []byte{'0', '0', '0', 'n', 'n'}
	require.False(t, blockchain.HashIsHardEnough(badData, 4))
	goodData := []byte{'0', '0', '0', '0', 'n'}
	require.True(t, blockchain.HashIsHardEnough(goodData, 4))
}

type fixedHash struct {
	hashResult []byte
}

func (f fixedHash) Write([]byte) (int, error) {
	return 0, nil
}

func (f fixedHash) Sum([]byte) []byte {
	return f.hashResult
}

func (f fixedHash) Reset() {}

func (f fixedHash) Size() int { return 0 }

func (f fixedHash) BlockSize() int { return 0 }

type miningTestHashFactory struct {
	calledGetHash int
	timesToFail   int
}

func (m *miningTestHashFactory) getHash() hash.Hash {
	m.calledGetHash++
	if m.calledGetHash <= m.timesToFail {
		return fixedHash{hashResult: []byte{'0', '0', '0', 'n', 'n'}}
	}
	return fixedHash{hashResult: []byte{'0', '0', '0', '0', 'n'}}
}

func Test_MineBlock_MeetsHardness(t *testing.T) {
	hashFactorySpy := miningTestHashFactory{timesToFail: 3}

	blockchain.MineBlock(nil, 0, "", 4, hashFactorySpy.getHash)

	require.Equal(t, 4, hashFactorySpy.calledGetHash)
}

func Test_MineBlock_FirstBlock(t *testing.T) {
	result, err := blockchain.MineBlock(nil, 0, "First Data", 1, sha256.New)

	require.Nil(t, err)
	fmt.Printf("MiningVariable: %d\n", result.MiningVariable)
}

func Test_MineBlock_Second(t *testing.T) {
	first, err := blockchain.MineBlock(nil, 0, "First Data", 1, sha256.New)
	require.Nil(t, err)

	second, err := blockchain.MineBlock(first.Hash, 1, "Second Data", 1, sha256.New)

	fmt.Printf("MiningVariable: %d\n", second.MiningVariable)
}
