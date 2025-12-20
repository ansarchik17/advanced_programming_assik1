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
