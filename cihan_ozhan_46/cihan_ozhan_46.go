package main

import "fmt"

func main() {
	sayHello("Hello", "Golang", "Developers", "Here") //Variadic Function is this one
	result := add(4,5,6,7,8,9,10)
	fmt.Println(result)
}

func sayHello(messages ...string) { //Variadic Function
	for _, message := range messages {
		fmt.Println(message)
	}
}

func add(terms ...int) int { //Variadic Function
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}
