package core

import (
	"log"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain  {
	println("new block chain")
	genesisBlock:=GenerateGenesisBlock()
	blockchain :=&Blockchain{}
	blockchain.appendBlcok(&genesisBlock)
	return blockchain
}

func (bc *Blockchain)Senddata(data string)  {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock :=GenerateNewBlcok(*preBlock,data)
	bc.appendBlcok(&newBlock)
}

func (bc *Blockchain) appendBlcok(newBlock *Block) {
	if len(bc.Blocks)==0{
		bc.Blocks =append(bc.Blocks,newBlock)
		return
	}
	if (isValidate(*newBlock, *bc.Blocks[len(bc.Blocks)-1])) {
		bc.Blocks = append(bc.Blocks, newBlock)
	}else {
		log.Fatal("invaid block")
	}
}

func isValidate(newBlock Block, oldBlock Block) bool {
	if (newBlock.Index != oldBlock.Index+1) {
		return false;
	}
	if (newBlock.PreBlockHash != oldBlock.Hash) {
		return false
	}
	if (newBlock.Hash != calculateHash(newBlock)) {
		println(newBlock.Hash)
		println(calculateHash(newBlock))
		return false
	}

	return true
}

func (bc *Blockchain) Print()  {
	for _,block :=range bc.Blocks{
		fmt.Printf("Index:%d\n",block.Index)
		fmt.Printf("Pre.has: %s\n",block.PreBlockHash)
		fmt.Printf("curr.has:%s\n",block.Hash)
		fmt.Printf("data:%s\n",block.Data)
		fmt.Printf("timestamp:%d",block.Timestamp)
		fmt.Println("\n")
	}
}
