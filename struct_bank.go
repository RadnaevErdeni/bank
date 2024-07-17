package bank

import (
	"sync"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID           int
	balance      float64
	transactions chan transaction
	wg           sync.WaitGroup
}

type transaction struct {
	operation string
	amount    float64
	response  chan interface{}
}

func NewAccount(id int) *Account {
	acc := &Account{
		ID:           id,
		transactions: make(chan transaction),
	}

	go acc.handleTransactions()

	return acc
}
