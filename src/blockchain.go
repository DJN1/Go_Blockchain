package src

// BlockChain type struct
type BlockChain struct {
	Blocks []*Block
}

// GenesisBlock returns the genesis Block
// @return Block: Genesis Block
func GenesisBlock() *Block {
	return NewBlock(-1, "", "Genesis Block")
}

// NewBlockChain creates a new Blockchain with genesis Block
// @return BlockChain: New Blockchain
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}

// getLastBlock returns the last Block in the Blockchain
// @return Block: Last Block of the Blockchain
func (bc *BlockChain) getLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks) - 1]
}

// validateBlock validates a Block
// @param newBlock Block: Block to validate
// @param previousBlock Block: Previous Block
// @return bool: True if Block is valid, false otherwise
func validateBlock(newBlock Block, previousBlock Block) bool {
	if((previousBlock.Index + 1) != newBlock.Index) {
		return false
	} else if(previousBlock.Hash != newBlock.PreviousHash) {
		return false
	} else {
		return true
	}
}

// addBlock adds a Block to the Blockchain
// @param data string: Data to add to the Block
func (bc *BlockChain) AddBlock(data string) {
	previousBlock := bc.getLastBlock()
	blockToAdd := NewBlock(previousBlock.Index, previousBlock.Hash, data)
	if(validateBlock(*blockToAdd, *previousBlock)) {
		bc.Blocks = append(bc.Blocks, blockToAdd)
	}
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		block.Print()
	}
}