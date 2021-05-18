package main

import "fmt"

/*func main() {
	fmt.Print("Please enter your Overall Grade: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println(value)
}*/

func main(){

	result, left := divide(104, 5)
	fmt.Println(result, left)

}


func divide (x, y int) (result, left int) {
	result = x/y
	left = x%y

	return result, left
}
