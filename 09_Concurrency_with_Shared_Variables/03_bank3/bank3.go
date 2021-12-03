package bank

import "sync"

var (
	balance      int
	balanceMutex sync.RWMutex
)

func Deposit(amount int) {
	balanceMutex.Lock()
	defer balanceMutex.Unlock()
	balance += amount
}

func Balance() int {
	balanceMutex.RLock()
	defer balanceMutex.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	balanceMutex.Lock()
	defer balanceMutex.Unlock()
	if balance < amount {
		return false
	}
	balance -= amount
	return true
}
