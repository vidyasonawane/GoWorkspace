package main

import "fmt"

//Everything in go is pass by value

func main() {
	x := 0
	foo(x)
	fmt.Println(x)
}

func foo(y int) {
	fmt.Println(y)
	y = 22
	fmt.Println(y)
}
