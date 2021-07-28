package BlockChain

import "test3/Block"
import "sync"

var BlockChain []Block.Block
var mutex sync.Mutex


func AddChain(block *Block.Block) []Block.Block {
	// 判断当前块的preHash和上一个的hash是否一致
	if isValidBlock(block) {
		mutex.Lock()
		BlockChain = append(BlockChain, *block)
		mutex.Unlock()
	}
	return BlockChain
}

func isValidBlock(block *Block.Block) bool  {
	if len(BlockChain) == 0 { // 创世区块
		return true
	}
	if block.PreHash != BlockChain[len(BlockChain)-1].MyHash {
		return false
	}

	return true
}

func Show() []Block.Block {
	return BlockChain
}


