package main

import "fmt"

func main() {
	fmt.Println((split(17)))

	numTerms, sum := add(1, 3)
	println("Added: ", numTerms, "Summary of Terms: ", sum)

}

func split(sum int) (x, y int) { //we name the return values here
	x = sum * 4 / 9 //no need to make "x:=" double dots here.
	y = sum - x
	return //so no need to show "x" and "y" in return area
}

func add (terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms) //the same double : dots are not needed
	return
}
