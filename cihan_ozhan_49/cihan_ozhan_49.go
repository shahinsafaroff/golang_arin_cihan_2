package main

import "fmt"

type human struct {
	FirstName string
	LastName string
	Age int
}

//Default and Empty "Method" creation function

func newHuman() *human {
	h := new(human)
	return h
}

func newHumanWithParams(firstName, lastName string, age int) *human {
	h := new(human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}

func main() {
	//Version 1 for struct realisation
/*	x := human{FirstName: "Shahin"}
	fmt.Println(x.FirstName)*/

	//Version 2 for struct realisation more used one!
/*	x := new(human)
	x.FirstName = "Shahin"
	fmt.Println(x.FirstName)*/

	//Creation method usage
/*	x := newHuman()
	x.Age = 28
	fmt.Println(x.Age)
*/
	x := newHumanWithParams("Muhammed", "Ali", 80)
	fmt.Println(x.Age)


}
