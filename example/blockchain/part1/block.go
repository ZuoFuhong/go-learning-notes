package main

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Timestamp     int64  // 区块时间戳
	PrevBlockHash string // 上一个区块的哈希值
	Hash          string // 当前区块的哈希值
	Data          string // 区块数据
}

// 计算区块哈希
func caculateHash(b *Block) string {
	blockData := string(b.Timestamp) + string(b.PrevBlockHash) + string(b.Data)
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// 生成区块
func GenerateNewBlock(preBlock *Block, data string) *Block {
	newBlock := &Block{}
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Data = data
	newBlock.Hash = caculateHash(newBlock)
	return newBlock
}

// 生成创世块
func GenerateGenesisBlock() *Block {
	preBlock := &Block{}
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
