package main

import "fmt"

type employeeStaff struct {
	 name      string
	 age       int
	 isMarried bool
	 kids []string
}

func main() {

	e1 := employeeStaff {
		name: "Nadav",
		age: 40,
		isMarried: false,
		kids: []string {
			"Arin",
			"Azazel",
		},
	}
	e2 := employeeStaff{
		name: "Shyamalan",
		age: 50,
		isMarried: true,
		kids: []string {
		"Memezok",
		"Pepezok",
		},
	}
	e3 := employeeStaff {
		name: "Sathan",
		age : 30,
		isMarried : true,
		kids: [] string {
			"Gabriel",
			"Mikael",
		},
	}
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
