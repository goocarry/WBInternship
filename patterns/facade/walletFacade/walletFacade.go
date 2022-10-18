package walletfacade

import (
	"fmt"

	account "github.com/goocarry/wb-internship/patterns/facade/account"
	billing "github.com/goocarry/wb-internship/patterns/facade/billing"
	securitycode "github.com/goocarry/wb-internship/patterns/facade/securityCode"
	wallet "github.com/goocarry/wb-internship/patterns/facade/wallet"
)

// WalletFacade ...
type WalletFacade struct {
	Account *account.Account
	Wallet *wallet.Wallet
	SecurityCode *securitycode.SecurityCode
	Billing *billing.Billing
}

// AddMoneyWithChecks ...
func (wf *WalletFacade) AddMoneyWithChecks(accountID string, amount int, code int) error {
	fmt.Println("Starting adding money")

	_, err := wf.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	wf.Wallet.AddMoney(amount)
	wf.Billing.SaveTransaction("123")
	return nil
}

// DeductMoneyWithChecks ...
func (wf *WalletFacade) DeductMoneyWithChecks(accountID string, amount int, code int) error {
	fmt.Println("Starting adding money")

	_, err := wf.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	_, err = wf.SecurityCode.CheckCode(code)
	if err != nil {
		return err
	}
	wf.Wallet.DeductMoney(amount)
	wf.Billing.SaveTransaction("321")
	return nil
}