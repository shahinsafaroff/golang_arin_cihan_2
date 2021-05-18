package main

import (
	"fmt"
	"math"
)
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
	area () float64
	circumference () float64
	diameter() float64
}
func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference())
	fmt.Println(i.diameter())
}
func main(){
	c1 := circle{5}
	interfaceFunc(c1)
}

