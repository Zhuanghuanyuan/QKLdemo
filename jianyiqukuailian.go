package main

import "strconv"
import "bytes"
import "crypto/sha256"
import "time"
import "fmt"

type Block struct {
	Timestamp     int64  //区块链被创建的时间
	Data          []byte //区块中实际包含的有用大的信息
	PrevBlockHash []byte //储存前一个区块的哈希值
	Hash          []byte //当前区块的哈希值
}

//区块链
type Blockchain struct {
	blocks []*Block
}

//写一个setHash()函数实现对SHA—256哈希的计算并且调用bytes包和sha256包中的函数
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//实现一个简化创建区块链的函数
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

//创建区块链：第一个创始块是固定的字符串内容；
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//用创始块创建一个区块链的函数
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//测试
func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
