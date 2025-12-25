package gym

import "fmt"

type BasicMember struct {
	Name  string
	Email string
}

type PremiumMember struct {
	Name        string
	Email       string
	BonusPoints int
}

type Member interface {
	GetName() string
	GetEmail() string
	GetPoints() int
}

type Gym struct {
	Members map[uint64]Member
}

func (basicMember BasicMember) GetName() string {
	return basicMember.Name
}

func (basicMember BasicMember) GetEmail() string {
	return basicMember.Email
}

func (basicMember BasicMember) GetPoints() int {
	return 0
}

func (premiumMember PremiumMember) GetName() string {
	return premiumMember.Name
}

func (premiumMember PremiumMember) GetEmail() string {
	return premiumMember.Email
}

func (premiumMember PremiumMember) GetPoints() int {
	return premiumMember.BonusPoints
}

func RunGymMenu(gym *Gym) {
	if gym.Members == nil {
		gym.Members = make(map[uint64]Member)
	}

	for {
		fmt.Println("\nGym Menu")
		fmt.Println("1. Show Members")
		fmt.Println("2. Add Member")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if len(gym.Members) == 0 {
				fmt.Println("No members in the gym")
			} else {
				fmt.Println("\nGym Members")
				for id, member := range gym.Members {
					fmt.Printf("ID: %d,  name: %s email: %s points: %d\n",
						id, member.GetName(), member.GetEmail(), member.GetPoints())
				}
			}

		case 2:
			var memberType string
			fmt.Print("Enter Member Type: ")
			fmt.Scanln(&memberType)

			var name, email string
			fmt.Print("Enter Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Email: ")
			fmt.Scanln(&email)

			var member Member
			if memberType == "Basic" || memberType == "basic" {
				member = BasicMember{Name: name, Email: email}
			} else if memberType == "Premium" || memberType == "premium" {
				var points int
				fmt.Print("Enter Bonus Points: ")
				fmt.Scanln(&points)
				member = PremiumMember{Name: name, Email: email, BonusPoints: points}
			} else {
				fmt.Println("We do not have this type")
				continue
			}

			var id uint64 = 1
			for {
				if _, exists := gym.Members[id]; !exists {
					break
				}
				id++
			}

			gym.Members[id] = member
			fmt.Println("Member added successfully!")

		case 3:
			fmt.Println("Exiting Gym Menu")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
