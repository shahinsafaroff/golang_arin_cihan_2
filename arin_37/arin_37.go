package main

import (
	"fmt"
)
//Arrays are underlying data type of Slices
//Pointers are the storage addresses of the data types

//Variables, Slices and Maps are "Pass by Reference" Data Types
// Arrays, Structs and functions are "Pass by Value" Data Types

func main() {

/*	name := "arin"

	fmt.Println(name)
	/*fmt.Println(&name) //& ---> address operator*/

/*	x := 22*/
/*	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println()
	fmt.Printf("%T, %\n", x, x)
	fmt.Printf("%T, %\n", &x, &x)*/

/*	 y := &x
	fmt.Printf("%T, %v\n", y, y)

	 z := &name
	fmt.Printf("%T, %v\n", z, z)
	 // * makes Deferencing which is changes address to value itself
	 fmt.Println(*(&x))

	fmt.Println(&(*(&x)))
	 //Again returns address of the value*/

/*	x1 := 10
	x2 := x1
	fmt.Println(x1, x2)
	x1 = 5
	fmt.Println(x1, x2)*/

/*	x1 := 10
	x2 := &x1

	fmt.Println(x1, x2)
	//*x2 grabs the value of the x1 via &x1
	fmt.Println(x1, *x2)

	*x2 = 3 //devaluing
	fmt.Println(x1, *x2)
	//*x2 matches &x1 place and then assigns "3" value to both x1 and x2

	x3 := &x1

	*x3 = 5

	fmt.Println(x1, *x2, *x3)
*/
/*	x1 := [4]int{1,10,100,1000}

	x2 := x1

	fmt.Println(x1, x2)

	x2[0] = 3
	fmt.Println(x2)
	fmt.Println(x1)
	//Array is "Pass by Value" Data type so the source array WON'T change*/
	x1 := []int{1,10,100,1000}

	x2 := x1

	fmt.Println(x1, x2)

	x2[0] = 3
	fmt.Println(x2)
	fmt.Println(x1)
	//Slice is "Pass by Reference" Data type so the source slice WILL change
}

