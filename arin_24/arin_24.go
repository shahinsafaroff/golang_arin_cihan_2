package main

import "fmt"

func main() {
	fmt.Println(result(49))

}


func result(grade int) string {
	if grade >= 50 {
		return "You've passed!"
	}
	return "You've failed!"
}


