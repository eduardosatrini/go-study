package bank

import (
	"fmt"
)

type Bitcoin float64

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.1f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(v Bitcoin) {
	w.balance += v
}

func (w *Wallet) Withdraw(v Bitcoin) {
	w.balance -= v
}
