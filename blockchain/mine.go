package blockchain

import (
	"hash"
	"math/rand"
	"time"
)

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

func MineBlockWithSleep(previousHash []byte, previousIndex int64, data string, minSeconds int, maxSeconds int, hashFactory HashFactory) (Block, error) {
	result, err := NewBlock(previousHash, previousIndex, rand.Int63(), data, hashFactory())
	if nil != err {
		return Block{}, err
	}
	secondsRange := maxSeconds - minSeconds
	time.Sleep(time.Duration(rand.Intn(secondsRange*1000))*time.Microsecond + time.Duration(minSeconds)*time.Second)
	return result, nil
}
