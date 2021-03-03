package blockchain

import (
	"hash"
)

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

func ChainIsValid(chain Blockchain, requiredLeadingZeros int, hashFactory HashFactory) bool {
	lastBlock := chain[0]
	if !FirstBlockIsValid(lastBlock, requiredLeadingZeros, hashFactory()) {
		return false
	}
	for i := 1; i < len(chain); i++ {
		if !IsValid(lastBlock, chain[i], requiredLeadingZeros, hashFactory()) {
			return false
		}
		lastBlock = chain[i]
	}
	return true
}

func getLongestChain(validChains []Blockchain) Blockchain {
	var longest Blockchain
	for _, chain := range validChains {
		if len(chain) > len(longest) {
			longest = chain
		}
	}
	return longest
}

func getValidBlockchains(chains []Blockchain, requiredLeadingZeros int, hashFactory HashFactory) []Blockchain {
	result := []Blockchain{}

	for _, chain := range chains {
		if ChainIsValid(chain, requiredLeadingZeros, hashFactory) {
			result = append(result, chain)
		}
	}

	return result
}

func GetMostValidBlockChain(chains []Blockchain, requiredLeadingZeros int, hashFactory HashFactory) Blockchain {
	validChains := getValidBlockchains(chains, requiredLeadingZeros, hashFactory)
	return getLongestChain(validChains)
}
