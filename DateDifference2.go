package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Date(2026, 12, 31, 23, 59, 0, 0, time.UTC)

    fmt.Println("Target Date: ",start)
    fmt.Println("Todays Date: ",time.Now())
    
        // calculate years, month, days and time betwen dates
        year, month, day,_,_,_ := diff(start, time.Now())

        fmt.Printf("Difference %d Years, %d Months, %d Days.\n\n", year, month, day)
        //fmt.Printf("Difference %d Years, %d Months, %d Days, %d Hours, %d Minutes and %d Seconds.\n\n", year, month, day, hour, min, sec)
		//fmt.Printf("")

        // calculate total number of days
		difference := start.Sub(time.Now())
		
    //fmt.Printf("Difference %d days", int(duration.Hours()/24) )
	fmt.Printf("Weeks: %d\n", int64(difference.Hours()/24/7))
   	fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
	fmt.Printf("Working Days: %d\n", int64(difference.Hours()/24) - 2 * int64(difference.Hours()/24/7))
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
    // if a.Location() != b.Location() {
    //     b = b.In(a.Location())
    // }
    if a.After(b) {
        a, b = b, a
    }
    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
    h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)
    hour = int(h2 - h1)
    min = int(m2 - m1)
    sec = int(s2 - s1)

    // Normalize negative values
    if sec < 0 {
        sec += 60
        min--
    }
    if min < 0 {
        min += 60
        hour--
    }
    if hour < 0 {
        hour += 24
        day--
    }
    if day < 0 {
        // days in month:
        t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
        day += 32 - t.Day()
        month--
    }
    if month < 0 {
        month += 12
        year--
    }

    return
}