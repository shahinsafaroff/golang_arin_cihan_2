package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	target := numRand(1, 100)

	fmt.Println("Find the number within 1 and 100")

	reader :=bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "tries left ")
		fmt.Print("Please enter guessed number: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("Your guess is higher, enter lesser number")
		} else if num < target {
			fmt.Println("Your guess is lesser, enter higher number")
		} else {
			fmt.Println("BINGO you've found number, try quantity was", target,  attempts, "th try found the number ")
			break
		}

	}

}
func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}