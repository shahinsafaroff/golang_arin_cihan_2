package main

import "fmt"

type employee struct {
	 name      string
	 age       int
	 isMarried bool
}
type manager struct {
	 employee
	 hasDegree bool
}
func main() {
	//via MAPS
	e1 := employee {
		name: "Nadav",
		age: 40,
		isMarried: true,
	}
//via variables
	e2 := employee{}
	e2.name = "Shyamalan"
	e2.age = 35
	e2.isMarried = false
//via MAPS
	m1 := manager {
		employee: employee{
			name: "Gerasim",
			age: 48,
			isMarried: false,
		}, hasDegree: true,
	}
//via Variables
	m2 := manager{}
	m2.name = "Rustam"
	m2.age = 32
	m2.isMarried = true
	m2.hasDegree = true

	//ANONIM STRUCT is used when the person is only one and such kind of struct won't be needed
	theBoss := struct {
		name string
		money bool
	}{name: "THE BOSS", money:true}

	fmt.Println(e2)
	fmt.Println(e1)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(theBoss)
}
