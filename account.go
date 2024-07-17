package bank

import "sync"

var (
	accounts = make(map[int]*Account)
	mu       sync.Mutex
	nextID   = 1
)

func CreateAccount() *Account {
	mu.Lock()
	defer mu.Unlock()

	account := NewAccount(nextID)
	accounts[nextID] = account
	nextID++
	return account
}

func GetAccount(id int) (*Account, bool) {
	mu.Lock()
	defer mu.Unlock()

	account, exists := accounts[id]
	return account, exists
}
