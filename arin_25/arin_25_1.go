package s

import (
	"errors"
	"fmt"
)

func main() {

	result, err := evenNum(20)
	if err  != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Entered number: ",  result)
	}

}

func evenNum (num int) (string, error){
	if num % 2 != 0 {
		return "", errors.New("There is an Error occured. You haven't entered EVEN Number")
	}
	return "You've entered EVEN NUMBER",  nil

}
