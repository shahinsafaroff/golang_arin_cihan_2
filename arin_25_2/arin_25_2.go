package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sRoot(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func sRoot(num float64) (float64, error){
	if num < 0 {
		return 0, errors.New("Negativ counts can't be entered for Square Root Operation")
	}

	return math.Sqrt(num), nil
}