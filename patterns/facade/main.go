package main

import (
	"github.com/goocarry/wb-internship/patterns/facade/account"
	billing "github.com/goocarry/wb-internship/patterns/facade/billing"
	securitycode "github.com/goocarry/wb-internship/patterns/facade/securityCode"
	"github.com/goocarry/wb-internship/patterns/facade/wallet"
	walletfacade "github.com/goocarry/wb-internship/patterns/facade/walletFacade"
)

func main() {
	wallet := walletfacade.WalletFacade{
		Account: account.NewAccount("account"),
		SecurityCode: securitycode.NewSecurityCode(777),
		Wallet: wallet.NewWallet(),
		Billing: billing.NewBilling(),
	}

	wallet.AddMoneyWithChecks("account", 777, 40)
	wallet.DeductMoneyWithChecks("account", 777, 20)
}
