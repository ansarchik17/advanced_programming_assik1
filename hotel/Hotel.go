package hotel

import "fmt"

type Room struct {
	RoomNumber    string
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Name  string
	Rooms map[string]Room
}

func (hotel *Hotel) AddRoom(room Room) {
	hotel.Rooms[room.RoomNumber] = room
	if hotel.Rooms == nil {
		hotel.Rooms = make(map[string]Room)
	}
}

func (hotel *Hotel) CheckIn(roomNumber string) bool {
	room, exists := hotel.Rooms[roomNumber]
	if !exists {
		fmt.Println("Room does not exist")
		return false
	}
	if room.IsOccupied {
		fmt.Println("Room already occupied")
		return false
	}
	room.IsOccupied = true
	hotel.Rooms[roomNumber] = room
	fmt.Println("Checked in successfully")
	return true
}

func (hotel *Hotel) CheckOut(roomNumber string) bool {
	room, exists := hotel.Rooms[roomNumber]
	if !exists {
		fmt.Println("Room does not exist")
		return false
	}
	if room.IsOccupied {
		fmt.Println("Room already occupied")
		return false
	}
	room.IsOccupied = false
	fmt.Println("Checked out successfully")
	return true
}

func (hotel *Hotel) ListVacantRooms() {
	fmt.Println("Hotel List of Vacant Rooms:")
	for number, room := range hotel.Rooms {
		if !room.IsOccupied {
			fmt.Printf("%s is occupied\n", number)
		}
	}
}

func (hotel *Hotel) GetRoomInfo(roomNumber string) (Room, bool) {
	room, exists := hotel.Rooms[roomNumber]
	return room, exists
}

func RunHotelMenu(myHotel *Hotel) {
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

			room := Room{
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
				continue
			}

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

		case 6:
			fmt.Println("Exiting Hotel menu")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
