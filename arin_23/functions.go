package main

import "fmt"

func main() {
/*	var x, y, sum int
	x = 5
	y = 10
	sum = x + y
	fmt.Printf("%d and %d summary is %d\n", x, y, sum )

	x = 7
	y = 11
	sum = x + y
	fmt.Printf("%d and %d summary is %d\n", x, y, sum )*/
	fmt.Println(sum(5, 10))

	hello("Shahin", 34)

}

func sum(x , y int) int {
	return x + y
}

func hello (name string, age int){
	fmt.Printf("My %s, age %d", name, age)
}

