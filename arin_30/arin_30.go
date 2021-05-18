package main

import "fmt"

func main() {

/*	mySlice := [] int {1,2,3}

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))

*/

	//Slices are dynamic arrays

/*	var mySlice []int
	mySlice = make([]int, 4)
	fmt.Println(mySlice)
	mySlice[0] = 10
	fmt.Println(mySlice)*/

	mySlice := []int{1,2,3}
	fmt.Println(mySlice)
	mySlice2 := mySlice
	fmt.Println(mySlice2)
	mySlice2[0] = 33
	fmt.Println(mySlice2)
	fmt.Println(mySlice)

}


