package main

import (
	"blockchain/code07-blc-print/BLC"
)

// 测试
func main()  {
	blockChain := BLC.CreatBlockchianWithGenesisBlock()
	defer blockChain.DB.Close()
	blockChain.AddBlock([]byte("szc"))
	blockChain.AddBlock([]byte("love"))
	blockChain.AddBlock([]byte("cx"))
	blockChain.PrintChain()
}
