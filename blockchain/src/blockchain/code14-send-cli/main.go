package main

import (
	"blockchain/code14-send-cli/BLC"

)

// 测试
func main()  {
	//blockChain := BLC.CreatBlockchianWithGenesisBlock()
	//defer blockChain.DB.Close()
	//blockChain.AddBlock([]byte("szc"))
	//blockChain.AddBlock([]byte("love"))
	//blockChain.AddBlock([]byte("cx"))
	cli := BLC.CLI{}
	cli.Run()
}
