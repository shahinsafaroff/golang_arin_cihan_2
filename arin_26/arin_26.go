package main

import "fmt"

func main () {
	x, y := 10, 4
	summ1, diff1, prod1 := calculation(x, y)
	fmt.Println("Summary: ", summ1)
	fmt.Println("Difference: ", diff1)
	fmt.Println("Production: ", prod1)

}

func calculation (num1, num2 int) (int, int, int){
	sum := num1 + num2
	diff := num1 - num2
	prod := num1 * num2

	return sum, diff, prod
}
