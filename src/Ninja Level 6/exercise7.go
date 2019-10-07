package main

import "fmt"

var g func()

func main() {
	f := func() {
		fmt.Println("Function expression")
	}

	f()
	fmt.Printf("%T\n", f)

	f2 := func(x int) {
		fmt.Println("Function expression with parameters", x)

	}
	f2(99)
	fmt.Printf("%T\n", f2)

	g = f
	g()
	fmt.Printf("%T\n", g)

}
