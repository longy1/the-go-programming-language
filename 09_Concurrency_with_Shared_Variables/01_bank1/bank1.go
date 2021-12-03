package bank

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)
var withdrawErr = make(chan error)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) error {
	withdraws <- amount
	return <-withdrawErr
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-withdraws:
			if balance-amount < 0 {
				withdrawErr <- fmt.Errorf("no enough money")
			} else {
				balance -= amount
				withdrawErr <- nil
			}
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
