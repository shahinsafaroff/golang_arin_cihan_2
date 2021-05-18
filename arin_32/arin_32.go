package main

import "fmt"

func main() {

/*	myMap := map[string]int {
		"Memezok": 40,
		"Pepezok": 17,
		"Meningit": 14,
		"Krisa": 70,
	}
	fmt.Println(myMap)
	fmt.Println(myMap["Memezok"], myMap["Selim"])*/

/*	isMarried := map[string] bool {
	"Memezok": true,
	"Pepezok": false,
	"Meningit": true,
	"Krisa": false,
	}
	fmt.Println(isMarried)*/

/*	myMap := make(map[string] int)
	fmt.Println(myMap)*/
	//creates empty map as slices were made via make method

	studentGrades := map[string]int {
		"Memezok": 80,
		"Pepezok": 47,
		"Meningit": 74,
		"Krisa": 0,
	}
/*	value, ok := studentGrades["Elis"]
	fmt.Println(value, ok)*/
//Returns value and avaliability
/*	value, ok := studentGrades["Memezok"]
	fmt.Println(value, ok)*/

/*	_, ok := studentGrades["Memezok"]
	fmt.Println(ok)*/
//Returns only availability

//Adding and Elements to a MAPs with alphabetical order

/*	fmt.Println(studentGrades)
	studentGrades["Sathan"] = 55
	fmt.Println(studentGrades)*/

//Deleting an elements from MAPS
/*	delete(studentGrades, "Meningit")
	fmt.Println(studentGrades)*/

//Length of MAPS

/*	fmt.Println(len(studentGrades))
//Creating new maps from other maps
	sg := studentGrades //maps are --> pass by reference as slices
	fmt.Println(sg)
	delete(sg, "Krisa")
	fmt.Println(sg)
	fmt.Println(studentGrades)*/

	for k, v := range studentGrades {
	fmt.Println(k, v)
}

}
