package main

import "fmt"

func main() {

	myArr := [10]int{0,1,2,3,4,5,6,7,8,9}

	myArr = mySquareArray(myArr) //First Class Arrays

	fmt.Println(myArr)


}

func mySquareArray(arr[10]int) [10]int {
	for i:=0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}