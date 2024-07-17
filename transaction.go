package bank

import (
	"bank/logger"
	"fmt"
)

func (a *Account) handleTransactions() {
	for t := range a.transactions {
		switch t.operation {
		case "deposit":
			a.balance += t.amount
			t.response <- nil
		case "withdraw":
			if a.balance >= t.amount {
				a.balance -= t.amount
				t.response <- nil
			} else {
				t.response <- fmt.Errorf("insufficient funds")
			}
		case "getBalance":
			t.response <- a.balance
		}
	}
}

func (a *Account) Deposit(amount float64) error {
	a.wg.Add(1)
	defer a.wg.Done()

	t := transaction{
		operation: "deposit",
		amount:    amount,
		response:  make(chan interface{}),
	}

	a.transactions <- t
	<-t.response
	logger.LogOperation("deposit", a.ID, amount)
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.wg.Add(1)
	defer a.wg.Done()

	t := transaction{
		operation: "withdraw",
		amount:    amount,
		response:  make(chan interface{}),
	}

	a.transactions <- t
	resp := <-t.response
	if resp != nil {
		return resp.(error)
	}
	logger.LogOperation("withdraw", a.ID, amount)
	return nil
}

func (a *Account) GetBalance() float64 {
	a.wg.Add(1)
	defer a.wg.Done()

	t := transaction{
		operation: "getBalance",
		response:  make(chan interface{}),
	}

	a.transactions <- t
	return (<-t.response).(float64)
}
