package model

import "strconv"

type Transaction struct {
	Sender    string
	Recepient string
	Value     int
	Signatue  string
}

func NewTransaction(transaction_fields []string) *Transaction {
	value, err := strconv.Atoi(transaction_fields[2])
	if err != nil {
		return nil
	}

	return &Transaction{
		Sender:    transaction_fields[0],
		Recepient: transaction_fields[1],
		Value:     value,
		Signatue:  transaction_fields[3],
	}
}

func (t *Transaction) ToData() string {
	return t.Sender + t.Recepient + strconv.Itoa(t.Value) + t.Signatue
}
