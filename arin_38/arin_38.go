/*package main

import "fmt"

//Functions are Pass by Value type Data Structures

func main() {
	x := 5
	fmt.Println(x)
	double(x)
	fmt.Println(x)
}

func double (num int) {
	num *=2
	fmt.Println(num)
}*/

/*package main

import "fmt"

//However due to Slices Functions are "Pass by Reference"

func main(){
	mySlice := []int{1,10,100}
	fmt.Println(mySlice)
	double(mySlice)
	fmt.Println(mySlice)
}

func double (slice []int) {
	for i:=0; i < len(slice); i++ {
		slice[i] *= 2
	}
	fmt.Println(slice)
}*/

/*package main

import "fmt"

//However due to Arrays Functions are "Pass by Value"

func main(){
	myArray := [3]int{1,10,100}
	fmt.Println(myArray)
	double(myArray)
	fmt.Println(myArray)
}

func double (array [3]int) {
	for i:=0; i < len(array); i++ {
		array[i] *= 2
	}
	fmt.Println(array)
}*/

package main

import "fmt"

//Functions are Pass by Value type Data Structures
//Pointers are used for changing bigger arrays slices or maps without copying them
//In other words Pointers find change the bigger datas in their own place without moving or copying them.

func main() {
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}

func double (num *int) { //pointer method
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}