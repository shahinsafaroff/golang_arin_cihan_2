package main

import (
	"fmt"
	"getcelcius"

)

func main() {
	fmt.Print("Please enter Celcius Degree: ")
	celcius, err := getcelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}

	kelvin := celcius + 273

	fmt.Println(celcius, "Celcius Degree is", kelvin, "Kelvin Degree")
	
}
