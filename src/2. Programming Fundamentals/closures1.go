package main

import "fmt"

var x = 0 //package level scope

func increment() int {
	x++ //block level scope
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}
