package main

import (
	"errors"
	"fmt"
)

func main() {

	myError := errors.New("This is an Error!")//Creating an error message for interfering info disclosure
	fmt.Println(myError.Error())

}
