package main

import "fmt"

func add(x, y int) int { //additional function which takes its parameters from main function
	return x + y
}

func main() { //main function uses additional functions
	fmt.Println(add(42, 13))
}
