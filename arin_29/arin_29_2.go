package main

import "fmt"

func main() {

	cities := [4]string{"istambul", "roma", "tehran", "belgrad"}

	for index, city := range cities {
		fmt.Println(index, city)
	}

}
