package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
	PrevBlock *Block
}

type Chain []Block //allows me have methods with slices of Blocks (Generally speaking i just find methods cleaner then functions)

var Blockchain Chain //Block chain in use by the system

//Stringer for println
func (block Block) String() string {
	/* Deprecated through Stringer integration
	fmt.Printf("\n")
	fmt.Printf("Index: %d\n", block.Index)
	fmt.Printf("Timestamp: %s\n", block.Timestamp)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %s\n", block.Hash)
	fmt.Printf("PrevHash: %s\n", block.PrevHash)
	fmt.Printf("\n") */

	return fmt.Sprintf("Index: %d\nTimestamp: %s\nData: %s\nHash: %s\nPrevHash: %s\n", block.Index,
		block.Timestamp, block.Data, block.Hash, block.PrevHash)
}

func (chain Chain) View() {
	var temp []string
	for i := 0; i < len(chain); i++ {
		temp = append(temp, fmt.Sprintf("[%d]", i))
		if i+1 != len(chain) {
			temp = append(temp, fmt.Sprintf("->"))
		}
	}
	fmt.Println(strings.Join(temp, " "))
}

func (block Block) toString() string {
	if block.PrevHash != "" { //does not equal nil
		return string(block.Index) + block.Timestamp + string(block.Data) + block.PrevHash
	} else {
		return string(block.Index) + block.Timestamp + string(block.Data)
	}
}

func (block *Block) CalcHash() string {
	var hashfactors string

	if block.PrevBlock == nil {
		hashfactors = block.toString()
	} else {
		hashfactors = block.toString() + block.PrevBlock.toString()
	}
	h := sha256.New()
	h.Write([]byte(hashfactors))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(Data string, oldBlock *Block) (Block, error) {
	var newBlock Block

	if oldBlock != nil {
		Log.Trace("Generating new block")
		newBlock.Index = oldBlock.Index + 1
		newBlock.PrevBlock = oldBlock
		newBlock.PrevHash = oldBlock.Hash
	} else {
		Log.Trace("Generating new Genesis block")
		newBlock.Index = 0
	}
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = Data
	newBlock.Hash = newBlock.CalcHash()

	return newBlock, nil
}

func (block *Block) IsBlockValid() bool {
	if block.PrevBlock != nil {
		if block.PrevBlock.Index+1 != block.Index {
			return false
		}
		if block.PrevBlock.Hash != block.PrevBlock.PrevHash {
			return false
		}
	}
	if block.CalcHash() != block.Hash {
		return false
	}
	return true
}

func (chain Chain) IsChainValid() bool {
	for i := 0; i < len(chain); i++ {
		if !chain[i].IsBlockValid() {
			return false
		}
	}
	return true
}

//for use when updating from online
func replaceChain(blockchain Chain) {
	if len(blockchain) > len(Blockchain) {
		Blockchain = blockchain
	}
}

//Update from the online server
func UpdateChain(http string) {

}
