package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Open("abc.rar")
	//Creating an error message for interfering info disclosure
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error : ", err)
	}
}
