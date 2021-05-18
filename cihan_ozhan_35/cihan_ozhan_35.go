package main

import "fmt"

func main() {
//flexible array
	myArray := [...]int{4,5,6,7,8}
	myArray[4] = 555
	fmt.Println(myArray)


}
