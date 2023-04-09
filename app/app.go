package app

import (
	"blockchain/app/model"
	"fmt"
	"time"
)

type App struct {
}

func (*App) Go() {
	chain := &model.Blockchain{
		Blocks: []*model.Block{{
			Index:       0,
			PrevHash:    "",
			Timestamp:   time.Now(),
			Transaction: model.NewTransaction([]string{"", "", "0", "Genesis block"}),
			Hash:        "",
		}},
		LastBlock: nil,
	}
	chain.LastBlock = chain.Blocks[0]
	ScanNewBlocks(chain)
	is_valid := chain.VerifyChain()
	fmt.Printf("\n\n\nChain is valid: %t\n\n\n", is_valid)
	chain.PrintChain()
}

func ScanNewBlocks(chain *model.Blockchain) {
	text := []string{"Sender address", "Recepient address", "Coin value", "Your signature"}
	var str string
	for {
		fmt.Println("New Transaction (Enter or Exit)")
		fmt.Scanln(&str)
		if str == "Exit" {
			return
		}
		var transaction_fields []string
		for _, out := range text {
			fmt.Println("Please, Enter " + out + ": ")
			fmt.Scanln(&str)
			transaction_fields = append(transaction_fields, str)
		}
		chain.AddBlock(model.NewTransaction(transaction_fields))
	}
}
