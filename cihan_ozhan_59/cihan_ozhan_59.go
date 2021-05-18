package main

import (
	"fmt"
	"time"
)

func main() {
	x := fmt.Println
	xTime := time.Date(2071, 10, 11, 20,00,20, 50, time.UTC)
	now := time.Now()
	x(now)
	x("------------")
/*	x(now.Year())
	x(now.Day())
	x(now.Date())
	x(now.Hour())
	x(now.Minute())
	x(now.Location())
	x(now.Weekday())*/
	//DATE COMPARISON before or after or equal
/*	x(xTime.Before(now))
	x(xTime.After(now))
	x(xTime.Equal(now))*/

	diff := now.Sub(xTime)
	x(diff)
	x(diff.Hours())

}
