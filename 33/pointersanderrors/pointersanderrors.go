// Package pointersanderrors demonstrates some important cases when using pointers and errors
//
// Deposit method uses a pointer receiver to allow the actual balance to be updated and not a copy in the function scope:
//
//	func (w *Wallet) Deposit(amount float64)
//
// Struct pointers allow cumbersome notation:
//
//	return (*w).balance
//
// To be automatically de-referenced using simpler syntax:
//
//	return w.balance
package pointersanderrors

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is the global error to be used when a balance does not allow withdrawal
var ErrInsufficientFunds = errors.New("insufficient funds")

// Bitcoin is our own type of float64 to improve readability and allow our own implementation of standard library methods
type Bitcoin float64

type Stringer interface {
	String() string
}

// Wallet contains the current (or state of) balance of a wallet
type Wallet struct {
	balance Bitcoin
}

// String allows our own definition on Bitcoin of %s of format strings in prints
func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

// Deposit adds the amount to the wallet at the address of the method receiver
//
// This method therefore updates the actual balance of wallet and not a copy of it
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the current balance of the wallet address of the method receiver
//
// A copy of the wallet could also be used with:
//
//	func (w Wallet) Balance() float64
//
// It is convention to match receiver types so, a pointer is used, to match Deposit:
//
//	func (w *Wallet) Deposit(amount float64)
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw subtracts the amount from the balance of provided method receiver wallet address
//
// Withdrawing more than is available returns the insufficient funds error
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
