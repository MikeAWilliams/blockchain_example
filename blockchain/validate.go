package blockchain

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

func IsValid(previous Block, block Block, requiredLeadingZeros int) bool {
	if previous.Index+1 != block.Index {
		return false
	}
	if !sliceEqual(previous.Hash, block.PreviousHash) {
		return false
	}
	if !HashIsHardEnough(block.Hash, requiredLeadingZeros) {
		return false
	}
	return true
}
