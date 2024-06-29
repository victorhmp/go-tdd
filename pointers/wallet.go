package pointers

import (
	"errors"
	"fmt"
)

// The Stringer interface is defined by the fmt package. It has one method:

// type Stringer interface {
// 	String() string
// }

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// Notice that I don't need to dereference the pointer here. Go does it for
	// me. A pointer to a struct is called a "struct pointer", and they are
	// automatically dereferenced.
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// VERY important to keep in mind:
// "In Go, when you call a function or a method the arguments are copied."
// So Go passes arguments by copy, not references. To get references, you need
// to use pointers.

// Big heads-up:
// Pointers can be nil. When a function returns a pointer to something, you
// need to make sure you check if it's nil or you might raise a runtime
// exception - the compiler won't help you here.
