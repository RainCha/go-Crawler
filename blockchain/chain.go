package block

// 区块链
type Blockchain struct {
	blocks []*Block // 数组中保存按顺序存放哈希值  map中保存hash → block的键值对
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
