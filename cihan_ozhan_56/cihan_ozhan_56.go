package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(8)

}
