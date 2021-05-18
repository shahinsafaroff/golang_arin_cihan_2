package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin) //catches input
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n') //gives value as a string to str variable
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
 // strconv.ParseFloat(strings.TrimSpace(str), 64) changes entered input to float64 and trims the unneded spaces
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", f)
	}

}
