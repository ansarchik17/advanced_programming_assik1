package main

import (
	"assignment1/employee"
	"assignment1/gym"
	"assignment1/hotel"
	"assignment1/wallet"
	"fmt"
)

func main() {
	//Gym
	myGym := gym.Gym{
		Members: make(map[uint64]gym.Member),
	}

	myGym.Members[1] = gym.BasicMember{Name: "Ali", Email: "ali@mail.com"}
	myGym.Members[2] = gym.PremiumMember{Name: "Dana", Email: "dana@mail.com", BonusPoints: 50}

	fmt.Println("   Gym Members   ")
	for id, member := range myGym.Members {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Points: %d\n",
			id, member.GetName(), member.GetEmail(), member.GetPoints())
	}

	//Employees
	employees := []employee.SalaryCalculator{
		employee.FullTime{MonthlySalary: 4000, BonusRate: 0.1},
		employee.PartTime{HourlyRate: 20, HoursWorked: 40},
		employee.Contractor{ProjectRate: 1000, ProjectsCompleted: 8},
		employee.Intern{DailyRate: 50, DaysWorked: 25},
	}

	fmt.Println("Employee")
	for i, emp := range employees {
		fmt.Printf("Employee %d salary: %.2f\n", i+1, emp.CalculateSalary())
	}

	//Hotel
	myHotel := hotel.Hotel{
		Name:  "My Hotel",
		Rooms: make(map[string]hotel.Room),
	}
	fmt.Println("Hotel management system:")
	hotel.RunHotelMenu(&myHotel)

	//Wallet
	myWallet := wallet.Wallet{
		Author:       "Ali",
		Balance:      0,
		Transactions: []float64{},
	}
	fmt.Println("Wallet management system:")
	wallet.RunWalletMenu(&myWallet)
}
