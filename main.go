package main

import (
	"BlockChain/blockchain"
	"fmt"
	//"strconv"
)

func main()  {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("sec Block")
	chain.AddBlock("tg Block")
	chain.AddBlock("ft Block")

	for _, block := range chain.Blocks{
		//fmt.Printf("POW:  %s\n", strconv.FormatBool(pow.Validate()) )
		fmt.Printf("POW:  %x\n", block.Hash )
		fmt.Println()
	}
}
