package main

import (
	"fmt"

	"github.com/AnsarKeles/Assignment1/employee"
	"github.com/AnsarKeles/Assignment1/gym"
	"github.com/AnsarKeles/Assignment1/hotel"
	"github.com/AnsarKeles/Assignment1/wallet"
)

func main() {
	myGym := gym.Gym{
		Members: make(map[uint64]gym.Member),
	}

	myGym.Members[1] = gym.BasicMember{Name: "Ali", Email: "ali@mail.com"}
	myGym.Members[2] = gym.PremiumMember{Name: "Dana", Email: "dana@mail.com", BonusPoints: 50}

	myHotel := hotel.Hotel{
		Name:  "My Hotel",
		Rooms: make(map[string]hotel.Room),
	}

	myWallet := wallet.Wallet{
		Author:       "Ali",
		Balance:      0,
		Transactions: []float64{},
	}

	employees := &employee.Employee{
		Intern:     []employee.Intern{},
		FullTime:   []employee.FullTime{},
		PartTime:   []employee.PartTime{},
		Contractor: []employee.Contractor{},
	}

	for {
		fmt.Println("\nMain Menu")
		fmt.Println("1. Gym Menu")
		fmt.Println("2. Employee Menu")
		fmt.Println("3. Hotel Menu")
		fmt.Println("4. Wallet Menu")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Gym menu")
			gym.RunGymMenu(&myGym)

		case 2:
			fmt.Println("Employee Menu")
			employee.RunEmployeeMenu(employees)

		case 3:
			fmt.Println("Hotel Menu")
			hotel.RunHotelMenu(&myHotel)

		case 4:
			fmt.Println("Wallet Menu")
			wallet.RunWalletMenu(&myWallet)

		case 5:
			fmt.Println("Exiting program")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
