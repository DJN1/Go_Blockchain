package main

import (
	src "github.com/DJN1/Go_Blockchain/src"
)

func main() {
	bc := src.NewBlockChain()
	bc.AddBlock("block 1")
	bc.AddBlock("block 2")
	bc.AddBlock("block 3")
	bc.AddBlock("block 4")
	bc.Print()
}
