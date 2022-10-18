package account

import "fmt"

// Account ...
type Account struct {
	name string
}

// NewAccount ...
func NewAccount(accountName string) *Account {
    return &Account{
        name: accountName,
    }
}

// CheckAccount ...
func (a *Account) CheckAccount(ID string) (bool, error) {
    fmt.Printf("Account %s is verified\n", ID)
	return true, nil
}