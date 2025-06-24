package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func newBlock(data string, PreBlockHash []byte) *block {
	newBlock := block{}
	newBlock.Data = []byte(data)
	newBlock.PrevBlockHash = PreBlockHash
	newBlock.Timestamp = time.Now().Unix()

	return newBlock.setHash()
}

func (block *block) setHash() *block {
	bytesTimestamp := []byte(strconv.FormatInt(block.Timestamp, 10))

	// combine the data to a new byte []byte
	headers := bytes.Join([][]byte{block.PrevBlockHash, []byte(block.Data), bytesTimestamp}, []byte{})

	hash := sha256.Sum256(headers)
	block.Hash = hash[:]

	return block
}

type blockChain struct {
	Blocks []*block
}

func NewBlockChain() *blockChain {
	return &blockChain{[]*block{&block{Data: []byte("first"), PrevBlockHash: []byte{}}}}
}

func (chain *blockChain) addBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := newBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func main() {
	chain := NewBlockChain()

	var cmd string

	for {
		fmt.Println("====================")
		fmt.Println("1. add a block")
		fmt.Println("2. show all blocks")
		fmt.Println("====================")

		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "1":
			fmt.Printf("input: ")

			scanner := bufio.NewScanner(os.Stdin)

			if scanner.Scan() {
				input := scanner.Text()
				chain.addBlock(strings.TrimSpace(input))
				chain.addBlock(strings.TrimSpace(string(input)))
			}
			fmt.Print("\033[H\033[2J")
		case "2":
			for num, block := range chain.Blocks {
				fmt.Printf("block[%d] = %s\n", num, strings.TrimSpace(string(block.Data)))
			}
		default:
			fmt.Println("exit")
		}
	}
}
