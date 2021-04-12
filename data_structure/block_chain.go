package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
)

type block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// 区块链的结构
func (thisBlock *block) setHash() *block {

	bytesTimestamp := []byte(strconv.FormatInt(thisBlock.Timestamp, 10))

	headers := bytes.Join([][]byte{thisBlock.PrevBlockHash, []byte(thisBlock.Data), bytesTimestamp}, []byte{})

	hash := sha256.Sum256(headers)

	thisBlock.Hash = hash[:]

	return thisBlock
}

//create a new block
func newBlock(data string, PrevBlockHash []byte) *block {
	newBlock := block{}
	newBlock.Data = []byte(data)
	newBlock.PrevBlockHash = PrevBlockHash
	newBlock.Timestamp = time.Now().Unix()

	return newBlock.setHash()
}

type blockChain struct {
	Blocks []*block
}

func (thisBlockChain *blockChain) addBlock(data string) {
	prevBlock := thisBlockChain.Blocks[len(thisBlockChain.Blocks)-1]
	newBlock := newBlock(data, prevBlock.Hash)
	thisBlockChain.Blocks = append(thisBlockChain.Blocks, newBlock)
}

func newGenesisBlock() *block {
	genesisBlock := block{}
	genesisBlock.Data = []byte("first block")
	genesisBlock.PrevBlockHash = []byte{}

	return &genesisBlock
}

func NewBlockChain() *blockChain {
	return &blockChain{[]*block{newGenesisBlock()}}
}

func main() {
	newBlockChain := NewBlockChain()

	var cmd string

	for {
		fmt.Println("\n====================")
		fmt.Println("please input the 1 to add the block into blockChain")
		fmt.Println("please input the 2 to show blockChain Data")
		fmt.Println("exit if your press any key")
		fmt.Println("====================")

		_, err := fmt.Scanf("%s", &cmd)

		if err != nil {
			fmt.Printf("have one error to show %s", err)
		}

		switch cmd {
		case "1":
			inputData := make([]byte, 1024)
			fmt.Println("\n====================")
			fmt.Printf("please input the the data in this block: ")

			_, err := os.Stdin.Read(inputData)
			if err != nil {
				fmt.Printf("have one error to show %s", err)
			}
			fmt.Println("====================")
			newBlockChain.addBlock(string(inputData))

		case "2":
			for number, block := range newBlockChain.Blocks {
				fmt.Println("\n====================")
				fmt.Println("======Printing the block chain======")
				fmt.Printf("number: %d block data is : %s", number, block.Data)
				fmt.Printf("current block hash is : %x \n", block.Hash)
				fmt.Printf("pre block hash is : %x \n", block.PrevBlockHash)
				fmt.Println("======Printing the block chain ======")
				fmt.Println("====================")
			}

		default:
			fmt.Println("exit successfully")
		}

	}

}
