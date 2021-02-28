package blockchain

import (
	"encoding/binary"
	"hash"
	"time"
)

type Block struct {
	PreviousHash   []byte
	CreatedTime    time.Time
	Index          int64
	MiningVariable int64
	Data           string
	Hash           []byte
}

func int64ToBytes(input int64) []byte {
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, uint64(input))
	return result
}

func (b Block) ComputeHash(hasher hash.Hash) ([]byte, error) {
	_, err := hasher.Write(b.PreviousHash)
	if nil != err {
		return nil, err
	}
	_, err = hasher.Write([]byte(b.CreatedTime.String()))
	if nil != err {
		return nil, err
	}
	_, err = hasher.Write([]byte(b.Data))
	if nil != err {
		return nil, err
	}
	_, err = hasher.Write(int64ToBytes(b.Index))
	if nil != err {
		return nil, err
	}
	_, err = hasher.Write(int64ToBytes(b.MiningVariable))
	if nil != err {
		return nil, err
	}
	hash := hasher.Sum(nil)
	return hash, nil
}

func NewBlock(previousHash []byte, previousIndex int64, miningVariable int64, data string, hasher hash.Hash) (Block, error) {
	result := Block{PreviousHash: previousHash, CreatedTime: time.Now(), Index: previousIndex + 1, MiningVariable: miningVariable, Data: data}
	hash, err := result.ComputeHash(hasher)
	if nil != err {
		return Block{}, err
	}
	result.Hash = hash
	return result, nil
}
