package main

import (
	"fmt"
	"os"
)

func main () {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Our file", file)
	}
}
