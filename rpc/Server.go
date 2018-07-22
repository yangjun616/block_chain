package main

import (
	"net/http"
	"block_chain/core"
	"io"
	"encoding/json"
)
var blockchain *core.Blockchain
func run()  {
	http.HandleFunc("/blockchain/get",blockchainGetHandler)
	http.HandleFunc("/blockchain/write",blockchainWritetHandler)
	http.ListenAndServe("localhost:8888",nil)

}

func blockchainGetHandler(w http.ResponseWriter,r *http.Request)  {
	blockchain.Print()
	bytes, error :=json.Marshal(blockchain)
	if error!=nil{
		http.Error(w,error.Error(),http.StatusInternalServerError)
	}
	io.WriteString(w,string(bytes))
}
func blockchainWritetHandler(w http.ResponseWriter,r *http.Request)  {
	blockData := r.URL.Query().Get("data")
	blockchain.Senddata(blockData)
	blockchainGetHandler(w,r)
}

func main() {
	blockchain =core.NewBlockchain()
	run()
}