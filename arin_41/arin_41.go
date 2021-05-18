package main

import "fmt"

type shape interface {
	area() float64
}
//... this is not spread operator!!!
//... This operator is variatic function parameter and is being used when we don't know the exact quantity of "shape" interface
func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println("Square: ", shape.area())
	}
}
type triangle struct {
	a float64
	h float64
}
 func (t triangle) area() float64 {
 	return (t.a * t.h) / 2
 }

 type square struct {
 	a float64
 }

 func (s square) area() float64{
 	return s.a * s.a
 }

 type rectangle struct {
 	a float64
 	b float64
 }

 func (r rectangle) area() float64{
 	return r.a * r.b
 }

func main() {
	t := triangle{3,4}
	s := square{4}
	r := rectangle{4,5}
	//Polymorphism is in under mentioned function
	printArea(t, s, r)
}
