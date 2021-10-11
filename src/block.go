package src

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block type struct
type Block struct {
	Index int
	Hash string
	PreviousHash string
	Data string
	Timestamp int64
}

// Generate a new block
// @param previousIndex - index of previous block
// @param previousHash string: previous block hash
// @param data string: data to be stored in the block
// @return Block: new block instance
func NewBlock(previousIndex int, previousHash string, data string) *Block {
	timestamp := time.Now().Unix()
	hashData := sha256.New()
	hashData.Write([]byte(fmt.Sprint(previousIndex) + previousHash + fmt.Sprint(timestamp) + data))
	hash := hex.EncodeToString(hashData.Sum(nil))
	return &Block{Index: previousIndex + 1, Hash: hash, PreviousHash: previousHash, Data: data, Timestamp: time.Now().Unix()}
}

// Print information about given block
func (b Block) Print() {
	fmt.Println("------------------Block------------------")
	fmt.Println("Index: ", b.Index)
	fmt.Println("Hash: ", b.Hash)
	fmt.Println("Previous Hash: ", b.PreviousHash)
	fmt.Println("Data: ", b.Data)
	fmt.Println("Timestamp: ", time.Unix(b.Timestamp, 0).Format(time.UnixDate))
	fmt.Println("-----------------------------------------")
}