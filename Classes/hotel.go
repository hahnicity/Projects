package main

import "fmt"
import "math/rand"


// Room Class //
type Room struct {
    dayAvailable, cost int
}

// Hotel Class //
type Hotel struct {
    emptyRegulars, emptyPenthouses []Room
    usedRegulars, usedPenthouses []Room
}

func (h *Hotel) ReserveRegularRoom(currentDay int) {
    if len(h.emptyRegulars) > 0 {
        h.emptyRegulars, h.usedRegulars = h.CycleRoom(h.emptyRegulars, h.usedRegulars)
        h.usedRegulars = h.PickRandomDayAvailable(currentDay, h.usedRegulars)
    } 
}

func (h *Hotel) ReservePenthouseRoom(currentDay int) {
    if len(h.emptyPenthouses) > 0 {
        h.emptyPenthouses, h.usedPenthouses = h.CycleRoom(h.emptyPenthouses, h.usedPenthouses)
        h.usedPenthouses = h.PickRandomDayAvailable(currentDay, h.usedPenthouses)
    }
}

func (h Hotel) PickRandomDayAvailable(currentDay int, used []Room) []Room {
    day := rand.Intn(5) + 1 + currentDay // Assume customer will stay from 1 to 5 days
    used[len(used)-1].dayAvailable = day
    return used
}

func (h Hotel) CycleRoom(before, after []Room) ([]Room, []Room) {
    roomToChoose := rand.Intn(len(before))
    before[roomToChoose] = before[len(before)-1]
    after = append(after, before[len(before)-1])
    before = before[0:len(before)-1]
    return before, after   
}

func (h *Hotel) CheckIfAvailable(currentDay int) {
    checkFunc := func(used, empty []Room) ([]Room, []Room){
        if len(used) > 0 {
            for i := 0; i < len(used); i++ {
                if used[i].dayAvailable >= currentDay {
                    used, empty = h.CycleRoom(used, empty)
                }           
            }
        }
        return used, empty
    }
    h.usedRegulars, h.emptyRegulars = checkFunc(h.usedRegulars, h.emptyRegulars)
    h.usedPenthouses, h.emptyPenthouses = checkFunc(h.usedPenthouses, h.emptyPenthouses)
}

// Hotel constructor //
func initializeRooms(h *Hotel, regularRooms, penthouseRooms int) *Hotel {
    h.emptyRegulars = make([]Room, 0)
    h.emptyPenthouses = make([]Room, 0)
    h.usedRegulars = make([]Room, 0)
    h.usedPenthouses = make([]Room, 0)
    for i := 0; i < regularRooms; i++ {
        h.emptyRegulars = append(h.emptyRegulars, Room{0, 200})
    }
    for i := 0; i < penthouseRooms; i++ {
        h.emptyPenthouses = append(h.emptyPenthouses, Room{0, 1000})
    }
    return h
}

func initializeHotel(regularRooms, penthouseRooms int) *Hotel {
    h := new(Hotel)
    h = initializeRooms(h, regularRooms, penthouseRooms)
    return h
}

// main
func main() {
    /*Cycle through a 365 day year with the Hotel simulation*/
    h := initializeHotel(200, 10)
    for day := 0; day < 365; day++ {
        customers := rand.Intn(21) //pick upwards of 20 customers a day
        for cust := 0; cust < customers; cust++ {
            // Give customers a %5 chance of picking a penthouse room
            if rand.Intn(100) <= 4 {
                h.ReservePenthouseRoom(day)
            } else {
                h.ReserveRegularRoom(day)    
            }
        }
        h.CheckIfAvailable(day)
        fmt.Println("EMPTY REGULARS", h.emptyRegulars)
        fmt.Println("EMPTY PENTHOUSES", h.emptyPenthouses)
        fmt.Println("USED REGULARS", h.usedRegulars)
        fmt.Println("USED PENTHOUSES", h.usedPenthouses)
    }
}
