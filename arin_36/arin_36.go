/*package main

import "fmt"

type mile float64
type navyMile float64
type kilometer float64
//mile data type is created so mile is Defined Type or Named Type
//NavyMile and Mile here are Defined and different data types so we can't apply mathematical applications on them
//even if both of them have underlying float64 data type.
func main() {
	var m mile
	m = 1.6
	fmt.Println(m)

	var nm navyMile
	nm = 1.852
	fmt.Println(nm)

//Second way of declaring  variables in go
	k :=kilometer(7.8)
	fmt.Println(k)

	fmt.Println(kilometer(m) + kilometer(nm) + k)
}
*/