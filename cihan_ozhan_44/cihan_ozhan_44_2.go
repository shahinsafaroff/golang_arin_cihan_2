package main

import "fmt"

func main() {
	numTerms, sum := add(3,4,5)
	fmt.Println("Added: ", numTerms, "terms for total of", sum)
}

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}