/*package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("abc.rar")
	//Creating an error message for interfering info disclosure
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		//Error logged
		os.Exit(1)
	}
}
*/