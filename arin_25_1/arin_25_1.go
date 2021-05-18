package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := evenNum(11)
	if err  != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Entered number: ",  result)
	}

}

func evenNum (num int) (int, error){
	if num % 2 != 0 {
		return 0, errors.New("There is an Error occured. You haven't entered EVEN Number")
	}
	return num,  nil

}
