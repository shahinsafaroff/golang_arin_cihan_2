package main

import (
	"fmt"
	"strconv"
)
func main() {
	var myString string = "1"
	var x int = 10
	var f float64  = 2.0
	fmt.Println(f)
	number, _ := strconv.Atoi(myString) //for converting string to integer
	fmt.Println(number)
	result := number + x
	var convert = int(f) //converts float64 to integer
	fmt.Println(result)
	fmt.Println("Result: " + strconv.Itoa(result)) //for converting integer to string
	fmt.Println(convert)
}
