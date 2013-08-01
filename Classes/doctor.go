package main

import "math/rand"

// Patient Class //
type Patient struct {
    timePerVisit int
    hasBeenSeen bool
}

// Doctor Class //
type Doctor struct {
    schedule map[int][]Patient
}

func (d *Doctor) ScheduleNewPatient(day int, p Patient) {
    /* Schedule a patient as soon as possible */
    for i := day; i < len(d.schedule); i++ {
        if len(d.schedule[i]) >= 16 {
            continue    
        } else {
            d.schedule[i] = append(d.schedule[i], p)   
            break
        }
    }
}

func (d *Doctor) SeeNextPatient(day int) {
    for p := range d.schedule[day] {
        if d.schedule[day][p].hasBeenSeen == false{
            d.schedule[day][p].hasBeenSeen = true
            d.schedule[day][p] = d.schedule[day][p]
        }    
    }
}

// main
func makeDoctor(daysToWork int) *Doctor {
    d := new(Doctor)
    d.schedule = make(map[int][]Patient)
    for i := 0; i < daysToWork; i++ {
        d.schedule[i] = make([]Patient, 0)   
    }
    return d
}

func main() {
    d := makeDoctor(5) // Make the doctor work 5 days
    day := 0
    for day < 5 {
        patients := rand.Intn(26) // Have up to 25 people call the office everyday
        for i := 0; i < patients; i++ {
            d.ScheduleNewPatient(day, Patient{30, false})
        }
        for i := 0; i < len(d.schedule[day]); i ++ {
            d.SeeNextPatient(day)    
        }
        day = day + 1
    }
}
