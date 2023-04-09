package model

import (
	"blockchain/app/helper"
	"time"
)

type Block struct {
	Index       int
	PrevHash    string
	Timestamp   time.Time
	Transaction *Transaction
	Hash        string
}

func newBlock(index int, prev_hash string, transaction *Transaction) *Block {
	generator := helper.HashGenerator{}
	timestamp := time.Now()
	hash := generator.Get(index, prev_hash, timestamp, transaction.ToData())
	return &Block{
		Index:       index,
		PrevHash:    prev_hash,
		Timestamp:   timestamp,
		Transaction: transaction,
		Hash:        hash,
	}
}
