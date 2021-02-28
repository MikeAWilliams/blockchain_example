package blockchain

import "hash"

func HashIsHardEnough(hash []byte, requiredLeadingZeros int) bool {
	if len(hash) < requiredLeadingZeros {
		return false
	}
	for i := 0; i < requiredLeadingZeros; i++ {
		if '0' != hash[i] {
			return false
		}
	}
	return true
}

type HashFactory func() hash.Hash

func MineBlock(previousHash []byte, previousIndex int64, data string, requiredLeadingZeros int, hashFactory HashFactory) (Block, error) {
	var miningVariable int64
	for {
		blockAttempt, err := NewBlock(previousHash, previousIndex, miningVariable, data, hashFactory())
		if nil != err {
			return Block{}, err
		}
		if HashIsHardEnough(blockAttempt.Hash, requiredLeadingZeros) {
			return blockAttempt, nil
		}
		miningVariable++
	}
}
