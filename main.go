package main

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash []byte
	Data []byte
	Prev []byte
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

func main() {

}
