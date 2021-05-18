package main

import "fmt"

func main() {

/*mySlc := []int{1,2,3}
fmt.Println(mySlc)*/

/*mySlc = append(mySlc, 4,5)
fmt.Println(mySlc)*/

/*mySlc2 := append(mySlc, 4,5)
fmt.Println(mySlc2)

mySlc[0] = 100

fmt.Println(mySlc)
fmt.Println(mySlc2)*/

/*	mySlc := []int{1,2,3}
	mySlc2 := []int{4,5}

	mySlc = append(mySlc, mySlc2...) //Spread operator is needed for getting rid of error. Spread operator splits the slice to its containing integers

	fmt.Println(mySlc)*/
/*
	mySlc := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(mySlc)
	mySlc = mySlc[2:] //deletes first 2 elements
	fmt.Println(mySlc)
	mySlc = mySlc[:len(mySlc) - 2] // deletes 2 elements from the end
	fmt.Println(mySlc)*/

/*	var myArr [4]int
	fmt.Println(myArr)

	var mySlc []int
	mySlc = make([]int, 4) //zero values
	fmt.Println(mySlc)

	var mySlc2 []bool
	mySlc2 = make([]bool, 2) //zero values
	fmt.Println(mySlc2)*/

	var mySlc3 []int
	fmt.Printf("%#v", mySlc3)

	fmt.Println()

	mySlc4 := make([]int, 0)
	fmt.Printf("%#v", mySlc4)

}
