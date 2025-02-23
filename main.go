package main

import (
	"fmt"

	"github.com/dharmini98/golang-blockchain/blockchain"
)

func main() {
	blockchain := blockchain.InitBlockChain()
	blockchain.AddBlock("This is the first data string")
	blockchain.AddBlock("This is the second data string")
	blockchain.AddBlock("This is the third data string")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.Prev)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
