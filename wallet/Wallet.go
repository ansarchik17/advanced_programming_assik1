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

func (wallet *Wallet) ShowTransactions() []float64 {
	wallet.Transactions = append(wallet.Transactions, wallet.Balance)
	return wallet.Transactions
}

func RunWalletMenu(wallet *Wallet) {
	for {
		fmt.Println("\nWallet Menu")
		fmt.Println("1. Add Money")
		fmt.Println("2. Spend Money")
		fmt.Println("3. Show Balance")
		fmt.Println("4. Show Transactions")
		fmt.Println("5. Exit")
		fmt.Print("Choose your option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter amount to add: ")
			fmt.Scanln(&amount)
			wallet.AddMoney(amount)

		case 2:
			var amount float64
			fmt.Print("Enter amount to spend: ")
			fmt.Scanln(&amount)
			wallet.SpendMoney(amount)

		case 3:
			wallet.GetBalance()

		case 4:
			wallet.ShowTransactions()

		case 5:
			fmt.Println("Exiting Wallet menu")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
