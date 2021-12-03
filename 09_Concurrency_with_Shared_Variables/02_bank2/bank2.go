package bank

var (
	sem     = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sem <- struct{}{}
	balance += amount
	<-sem
}

func Balance() int {
	sem <- struct{}{}
	defer func() { <-sem }()
	return balance
}
