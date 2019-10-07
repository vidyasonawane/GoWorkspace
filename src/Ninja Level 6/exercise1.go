package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	p, q := bar()
	fmt.Println(p)
	fmt.Println(q)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	a := 22
	b := "vidya"
	return a, b
}
