package main

import (
	"fmt"
	"strconv"

	"github.com/Danvaspri/golang-blockchain/blockchain"
)


func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("new block")
	chain.AddBlock("second after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.Blocks {

		fmt.Printf("Prev Hash: %x\n",block.PrevHash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Hash:\r%x\n",block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}


}