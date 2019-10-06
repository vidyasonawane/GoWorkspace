package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("My first function expression")
	}

	f()

	f2(22)
}

var f2 = func(x int) {
	fmt.Println("Function expression with number", x)
}
