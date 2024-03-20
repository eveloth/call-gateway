package models

import (
	"time"

	money "github.com/Rhymond/go-money"
)

type Operation string

const (
	INC Operation = "INC"
	DEC Operation = "DEC"
)

type Transaction struct {
	Id        int
	UserId    int
	Operation Operation
	Amount    money.Amount
	Timestamp time.Time
}

func NewTxInc(userId int, amount money.Amount, ts time.Time) *Transaction {
	return &Transaction{
		UserId:    userId,
		Operation: INC,
		Amount:    amount,
		Timestamp: ts,
	}
}

func NewTxDec(userId int, amount money.Amount, ts time.Time) *Transaction {
	return &Transaction{
		UserId:    userId,
		Operation: DEC,
		Amount:    amount,
		Timestamp: ts,
	}
}
