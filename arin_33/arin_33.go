package arin_33

import "fmt"

/*func main() {
	//Array

	myArr := [5]string{"Tel-Aviv", "Belgrad", "Roma", "Tbilisi", "Moscow"}

	for index, value := range myArr {
		fmt.Println(index, value)
	}

	fmt.Println()
	// Slices

	mySlice := []int{1,2,3,4,5}

	for i, v  := range mySlice {
		fmt.Println(i, v)

	}*/

/* func main() {
 	mySlice := []int {0,1,2,3,4,5,6,7,8,9}

 	mySlice1 := append(mySlice[2:4], mySlice[6:8])

 	fmt.Println(mySlice1)

}*/

//Slices copy method vs assign

/*func main() {
	mySlice := []int{1,2,3}
	mySlice2 := make([]int,2)

	fmt.Println(mySlice)
	fmt.Println(mySlice2)

	copy(mySlice2,mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice2)

	mySlice[0]= 100
	fmt.Println(mySlice)
	fmt.Println(mySlice2)
}*/

/*func main() {
	mySlice := []int{1,2,3}
	mySlice2 := make([]int,2)

	fmt.Println(mySlice)
	fmt.Println(mySlice2)

	mySlice2 = mySlice
	fmt.Println(mySlice)
	fmt.Println(mySlice2)

	mySlice[0]= 100
	fmt.Println(mySlice)
	fmt.Println(mySlice2)
}
*/

func main() {
	myMap := map[string][]string {
		"Emmanuel Nadav":[]string{"Aristokrat Rotshild", "Iron Fist"},
		"Azakiel Sauron":[]string{"Ugly Turks", "Lion in Cage", "Sauron"},
		"Nadav Essehrman":[]string{"Arabian idiots", "Turkish fantasy", "Sathan's sons" },
	}

	for key, value := range  myMap {
		fmt.Println("Writer", key)
		fmt.Println("Some books: ")
		for i, v := range value {
			fmt.Println("\t",i+1,v)
		}
	}

}