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

hotelLoop:
	for {
		fmt.Println("\nHotel Management System")
		fmt.Println("1. Add Room")
		fmt.Println("2. Check In")
		fmt.Println("3. Check Out")
		fmt.Println("4. List Vacant Rooms")
		fmt.Println("5. Get Room Info")
		fmt.Println("6. Exit")
		fmt.Print("Choose your option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var number, roomType string
			var price float64
			fmt.Print("Enter room number: ")
			fmt.Scanln(&number)
			fmt.Print("Enter room type: ")
			fmt.Scanln(&roomType)
			fmt.Print("Enter price per night: ")
			fmt.Scanln(&price)

			room := hotel.Room{
				RoomNumber:    number,
				Type:          roomType,
				PricePerNight: price,
				IsOccupied:    false,
			}
			myHotel.AddRoom(room)
			fmt.Println("Room added successfully.")

		case 2:
			var number string
			fmt.Print("Enter room number to check in: ")
			fmt.Scanln(&number)
			myHotel.CheckIn(number)

		case 3:
			var number string
			fmt.Print("Enter room number to check out: ")
			fmt.Scanln(&number)
			myHotel.CheckOut(number)

		case 4:
			myHotel.ListVacantRooms()

		case 5:
			var number string
			fmt.Print("Enter room number: ")
			fmt.Scanln(&number)
			room, exists := myHotel.GetRoomInfo(number)
			if !exists {
				fmt.Println("Room does not exist")
			} else {
				status := "Vacant"
				if room.IsOccupied {
					status = "Occupied"
				}
				fmt.Printf(
					"Room %s, Type: %s, Price: %.2f, Status: %s\n",
					room.RoomNumber,
					room.Type,
					room.PricePerNight,
					status,
				)
			}

		case 6:
			fmt.Println("Exiting Hotel menu")
			break hotelLoop

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}

	//Wallet
	myWallet := wallet.Wallet{
		Author:       "Ali",
		Balance:      0,
		Transactions: []float64{},
	}

walletLoop:
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
			myWallet.AddMoney(amount)
			myWallet.GetTransactions()

		case 2:
			var amount float64
			fmt.Print("Enter amount to spend: ")
			fmt.Scanln(&amount)
			myWallet.SpendMoney(amount)
			myWallet.GetTransactions()

		case 3:
			myWallet.GetBalance()

		case 4:
			fmt.Println("Transaction history:")
			for i, t := range myWallet.Transactions {
				fmt.Printf("%d: %.2f\n", i+1, t)
			}

		case 5:
			fmt.Println("Exiting Wallet menu")
			break walletLoop

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}

	fmt.Println("\nProgram ended successfully!")
}
