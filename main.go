package main

import "fmt"

func main() {
	y := simpleFunction(1)
	fmt.Println(y)
}

func simpleFunction(x int) int {
	return x * x
}
