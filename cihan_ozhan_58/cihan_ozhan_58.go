package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2020, time.February, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Works: %s\n", t)

	fmt.Println("************")
	now := time.Now()
	fmt.Println("Current Time Date: %s\n", now)
	fmt.Println("************")
	fmt.Println("Year: ", now.Year())
	fmt.Println("************")
	fmt.Println("Month: ", now.Month())
	fmt.Println("************")
	fmt.Println("WeekDay: ", now.Weekday())
	fmt.Println("************")
	fmt.Println("Day of the Month: ", now.Day())
	tomorrow := now.AddDate(0,0,1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())
	fmt.Println("************")
	longFormat := "Monday, January 2, 2021"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
	fmt.Println("************")
	shortFormat := "1/2/2022"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
