package main

import(
	"fmt"
	"time"
)

func main(){
	firstDate := time.Date(2026, 12, 31, 24, 0, 0, 0, time.UTC)
   	//secondDate := time.Date(20, 2, 12, 6, 0, 0, 0, time.UTC)
	secondDate := time.Date(time.Now().Year(),time.Now().Month(),time.Now().Day(),time.Now().Hour(),time.Now().Minute(),0,0,time.UTC)
   	difference := firstDate.Sub(secondDate)
  	fmt.Println("The difference between dates", firstDate, "and", secondDate, "is: ")
   	fmt.Printf("Years: %d\n", int64(difference.Hours()/24/365))
   	fmt.Printf("Months: %d\n", int64(difference.Hours()/24/30)%12)
	fmt.Printf("Days: %d\n", int64(difference.Hours()/24)%30)

   	fmt.Printf("Weeks: %d\n", int64(difference.Hours()/24/7))
   	fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
	fmt.Printf("Working Days: %d\n", int64(difference.Hours()/24) - 2 * int64(difference.Hours()/24/7))
   	//fmt.Printf("Hours: %.f\n", difference.Hours())
   	//fmt.Printf("Minutes: %.f\n", difference.Minutes())
   	//fmt.Printf("Seconds: %.f\n", difference.Seconds())
   	//fmt.Printf("Nanoseconds: %d\n", difference.Nanoseconds())
}