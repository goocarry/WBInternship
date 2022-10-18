package wallet

import "fmt"

// Wallet ...
type Wallet struct {
	balance int
}

// NewWallet ...
func NewWallet() *Wallet {
    return &Wallet{
        balance: 0,
    }
}

// AddMoney ...
func (w Wallet) AddMoney(amount int)  {
	w.balance += amount
    fmt.Println("Wallet balance added successfully")
    return
}

// DeductMoney ...
func (w Wallet) DeductMoney(amount int) error {
	if w.balance < amount {
        return fmt.Errorf("Balance is not sufficient")
    }
    fmt.Println("Wallet balance is Sufficient")
    w.balance = w.balance - amount
    return nil
}