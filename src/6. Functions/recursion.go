package main

import "fmt"

func main() {
	x := factorial(4)
	fmt.Println(x)

	y := loopfact(4)
	fmt.Println(y)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopfact(n int) int {
	i := 1
	for ; n > 0; n-- {
		i *= n
	}
	return i
}
