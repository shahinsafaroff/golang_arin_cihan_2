package main

import "fmt"

func main() {
//Anonymous functions are used only once so they do not needed to be called several times
	addFunc := func(terms ...int) (numTerms int, sum int){ //Anonymous Function
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}
	numTerms, sum := addFunc(2,5)
	fmt.Println("Added: ", numTerms, "Total summary of Terms: ", sum)
}
