package blockchain

import "hash"

func sliceEqual(one []byte, two []byte) bool {
	if len(one) != len(two) {
		return false
	}
	for index, bOne := range one {
		if bOne != two[index] {
			return false
		}
	}
	return true
}

func IsValid(previous Block, block Block, requiredLeadingZeros int, hasher hash.Hash) bool {
	if previous.Index+1 != block.Index {
		return false
	}
	if !sliceEqual(previous.Hash, block.PreviousHash) {
		return false
	}
	if !HashIsHardEnough(block.Hash, requiredLeadingZeros) {
		return false
	}
	expectedHash, err := block.ComputeHash(hasher)
	if nil != err {
		return false
	}
	if !sliceEqual(expectedHash, block.Hash) {
		return false
	}
	return true
}

func FirstBlockIsValid(block Block, requiredLeadingZeros int, hasher hash.Hash) bool {
	if 0 != block.Index {
		return false
	}
	if nil != block.PreviousHash {
		return false
	}
	if !HashIsHardEnough(block.Hash, requiredLeadingZeros) {
		return false
	}
	expectedHash, err := block.ComputeHash(hasher)
	if nil != err {
		return false
	}
	if !sliceEqual(expectedHash, block.Hash) {
		return false
	}
	return true
}
