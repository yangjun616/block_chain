package main

import "block_chain/core"

func main() {
	bc:=core.NewBlockchain()
	bc.Senddata("send 1 btc to jack")
	bc.Senddata("send 1 btc to lucy")
	bc.Print()
}

