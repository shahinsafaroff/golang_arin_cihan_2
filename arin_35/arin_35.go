/*package main

import "fmt"

//Structs are similar to Arrays and have "Pass by Value" behaviour
//However Maps and Slices are "Pass by Reference"
// Classes have an inheritance however structs are independent
type employee struct {
	name      string
	age       int
	isMarried bool
}
	func main() {
		e1 := employee {
			name: "Nadav",
			age: 40,
			isMarried: true,
		}

		fmt.Println(e1)
		e2 := e1
		fmt.Println(e2)
		e2.name = "Azakiel"
		fmt.Println(e2)
		fmt.Println(e1)

}
*/