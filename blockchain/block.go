package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// 区块结构
type Block struct {
	Timestamp     int64  //当前的时间戳（该区块何时创建）
	Data          []byte // 在区块中保存的有价值的信息
	PrevBlockHash []byte // 上一区块的散列值
	Hash          []byte // 该区块的散列值
}

// 采用一种简单的算法计算散列值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// 简化区块的创建
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 区块链
type Blockchain struct {
	blocks []*Block // 数组中保存按顺序存放哈希值  map中保存hash → block的键值对
}

// 实现像链中添加区块的方法
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// 创建创世块函数
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// 实现一个用创世区块创建区块链的函数
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// 定义挖矿难度
const targetBits = 24

// 工作量证明结构
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// 新建一个工作量证明
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

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
