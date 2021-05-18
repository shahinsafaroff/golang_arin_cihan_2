package main

import (
	"fmt"
	"math/rand"
	"strings"
	"github.com/shahinsafaroff/color"
	"../myPackage"
)

func main() {
	fmt.Println("My favorite number is", rand.Int63n(1000000000000002))
	//random integer generator for session ids
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.HasPrefix("unit_test", "unit")) // Here "unit_" is prefix
	fmt.Println(strings.HasSuffix("unit_test", "test")) // Here "_test" is suffix
	fmt.Println(strings.Index("test", "s"))
}
