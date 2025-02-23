package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash []byte
	Data []byte
	Prev []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Prev}, []byte{})
	hash := sha256.Sum256(info) //insecure hashing function, but this will do for now
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	new := CreateBlock(data, chain.Blocks[len(chain.Blocks)-1].Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block { //creation of the first block
	new := CreateBlock("Start", []byte{})
	return new //returns the first block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
