package Block

import (
	"fmt"
	"github.com/brady-wang/go-tools/hashx"
	"strconv"
	"strings"
	"time"
)

const difficult = 4


type Block struct {
	Index int
	BMP int
	Nonce int
	MyHash string
	PreHash string
	TimeStamp string
	Difficult int
}


// CreateFirstBlock 创世区块
func CreateFirstBlock() *Block {
	var block = Block{
		Index:     0,
		BMP:       0,
		Nonce:     0,
		PreHash:   "",
		TimeStamp: time.Now().String(),
		Difficult: difficult,
	}
	block.MyHash = CalculateHash(&block)

	return &block
}

// NextBlock 下一个区块
func NextBlock(BMP int,preBlock Block) *Block{
	var block = Block{
		Index:     preBlock.Index+1,
		BMP:       BMP,
		PreHash:   preBlock.MyHash,
		Nonce: preBlock.Nonce,
		TimeStamp: time.Now().String(),
		Difficult: difficult,
	}
	for {
		block.Nonce++
		hash := CalculateHash(&block)
		fmt.Println(hash)
		if strings.HasPrefix(hash,strings.Repeat("0",block.Difficult)) {
			block.MyHash = hash
			fmt.Println("挖矿成功 " + hash)
			break
		}
	}
	return &block
}

// CalculateHash 计算hash
func CalculateHash(block *Block) string  {
	hashData := strconv.Itoa(block.Index) + strconv.Itoa(block.Nonce)+block.PreHash
	hash := hashx.Sha256(hashData)
	return hash
}
