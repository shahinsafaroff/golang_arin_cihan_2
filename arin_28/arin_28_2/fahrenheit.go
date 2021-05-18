package main

import (
	"fmt"
	"getcelcius"

)

func main() {
	fmt.Print("Please enter your Celcius degree: ")
	celcius, err := getcelcius.GetCelcius()
	if err !=  nil {
		fmt.Println(err)
	}

	fahrenheit := (celcius * 9/5) + 32

	fmt.Println(celcius, " Celcius degree is", fahrenheit, "equal Fahrenheit Degree.")

}
