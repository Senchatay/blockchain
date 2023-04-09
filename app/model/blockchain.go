package model

import (
	"fmt"
)

type Blockchain struct {
	Blocks    []*Block
	LastBlock *Block
}

func (chain *Blockchain) AddBlock(transaction *Transaction) {
	lastBlock := chain.LastBlock
	newBlock := newBlock(lastBlock.Index+1, lastBlock.Hash, transaction)

	chain.Blocks = append(chain.Blocks, newBlock)
	chain.LastBlock = newBlock
}

func (chain *Blockchain) VerifyChain() bool {
	for i := 1; i < len(chain.Blocks); i++ {
		if chain.Blocks[i].PrevHash != chain.Blocks[i-1].Hash {
			return false
		}
	}
	return true
}

func (chain *Blockchain) PrintChain() {
	for _, block := range chain.Blocks {
		fmt.Printf(
			"Index: %d\nPrevHash: %s\nTimestamp: %s\nData: %s\nHash: %s\n\n",
			block.Index,
			block.PrevHash,
			block.Timestamp.String(),
			block.Transaction.ToData(),
			block.Hash,
		)
	}
}
