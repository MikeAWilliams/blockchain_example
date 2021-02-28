package blockchain

func HashIsHardEnough(hash []byte, requiredLeadingZero int) bool {
	if len(hash) < requiredLeadingZero {
		return false
	}
	for i := 0; i < requiredLeadingZero; i++ {
		if '0' != hash[i] {
			return false
		}
	}
	return true
}
