package blockchain

func IsValid(previous Block, block Block, requiredLeadingZeros int) bool {
	if previous.Index+1 != block.Index {
		return false
	}
	return true
}
