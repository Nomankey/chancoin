package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)



type block struct {
	Data	string
	Hash	string
	PrevHash	string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block)getHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
} 

func getPrevHash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocks-1].Hash
}


func createBlock(Data string) *block {
	newBlock := block{Data, "", getPrevHash()}
	newBlock.getHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}

