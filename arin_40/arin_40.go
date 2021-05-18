/*package main

import (
	"fmt"
	"math"
)
type rectangle struct {
	a, b float64
}
func (r rectangle) area() float64 {
	return r.a * r.b
}
func (r rectangle) circumference() float64{
	return 2*(r.a + r.b)
}
//Interface is used to create methods for further usage
//Interfaces are always used together with Structs
type circle struct {
	r float64
}
func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}
func (c circle) diameter() float64 {
	return 2 * c.r
}
//Interface is used to create methods for further usage
type shape interface {
	area() float64
	circumference() float64
}
func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference())
	fmt.Printf("%T", i)
	fmt.Println()
}
func main(){
	r1 := rectangle{3,8}
	interfaceFunc(r1)
	fmt.Println("Square is: ", r1.area())
	fmt.Println("Perimeter is: ", r1.circumference())

	interfaceFunc(r1)
}

*/