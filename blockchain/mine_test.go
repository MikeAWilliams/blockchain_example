package blockchain_test

import (
	"blockchain_example/blockchain"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HashIsHardEnough(t *testing.T) {
	badData := []byte{'0', '0', '0', 'n', 'n'}
	require.False(t, blockchain.HashIsHardEnough(badData, 4))
	goodData := []byte{'0', '0', '0', '0', 'n'}
	require.True(t, blockchain.HashIsHardEnough(goodData, 4))

}
