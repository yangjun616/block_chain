package gotest

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

type hash_demo struct {

}

func calculateHash(toBeHashed string) string  {
	// go实现 md5 sha256 512
	//sha256hash 场地为256位

	hashInBytes :=sha256.Sum256([]byte(toBeHashed))

	hashInString := hex.EncodeToString(hashInBytes[:])
	log.Printf("%s,%s",toBeHashed,hashInString)

	return hashInString
}


/*func main()  {

	calculateHash("test1")
	calculateHash("test1")
	calculateHash("test2")
}*/
