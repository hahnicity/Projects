package main

import "rand"

// Room Class //
type Room struct {
    dayAvailable, cost int
}

// Hotel Class //
type Hotel struct {
    availableRegularRooms, availablePenthouseRooms []Room
    usedRegularRooms, usedPenthouseRooms []Room
}

func (hotel *Hotel) ReserveRegularRoom() {
    
    
}

func (hotel *Hotel) ReservePenthouseRoom() {
    

}

// Hotel constructor //
func initializeRooms(hotel Hotel, regularRooms, penthouseRooms int) *Hotel {
    hotel.availableRegularRooms = make([]Room, 0)
    hotel.availablePenthouseRooms = make([]Room, 0)
    hotel.usedRegularRooms = make([]Room, 0)
    hotel.usedPenthouseRooms = make([]Room, 0)
    for i := 0; i < regularRooms; i++ {
        hotel.availableRegularRooms = Room{0, 200}
    }
    for i := 0; i < penthouseRooms; i++ {
        hotel.availablePenthouseRooms = Room{0, 1000}
    }
    return hotel
}

func initializeHotel(regularRooms, penthouseRooms int) *Hotel {
    hotel := new(Hotel)
    hotel = initializeRooms(hotel, regularRooms, penthouseRooms)
    return hotel
}

// main
func main() {
    
    
}
