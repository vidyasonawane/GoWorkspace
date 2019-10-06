package main

import "fmt"

func main() {
	s1 := foo() //returning a string
	fmt.Println(s1)

	x := bar() //returning a function

	fmt.Printf("%T\n", x)

	fmt.Println(x)
	i := x()
	fmt.Println(i)
}

func foo() string {
	return "Vidya"
}

func bar() func() int {
	return func() int {
		return 2019
	}
}
