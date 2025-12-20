package wallet

import "fmt"

type Wallet struct {
	Author       string
	Balance      float64
	Transactions []float64
}

func (wallet *Wallet) AddMoney(amount float64) {
	wallet.Balance += amount
	fmt.Println(wallet.Balance)
}

func (wallet *Wallet) SpendMoney(amount float64) {
	if amount > wallet.Balance {
		fmt.Println("you do not have enough money")
		return
	}
	wallet.Balance -= amount
	fmt.Printf("You spent %.2f\n money, and your total balance is %.2f\n", amount, wallet.Balance)
}

func (wallet *Wallet) GetBalance() {
	fmt.Printf("Wallet balance: %.2f\n", wallet.Balance)
}

func (wallet *Wallet) GetTransactions() []float64 {
	wallet.Transactions = append(wallet.Transactions, wallet.Balance)
	return wallet.Transactions
}
