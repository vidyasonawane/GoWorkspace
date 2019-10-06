package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(xi...)
	fmt.Println("Sum of all numbers :", total)

	eventotal := even(sum, xi...)
	fmt.Println("Sum of even numbers: ", eventotal)

	oddtotal := odd(sum, xi...)
	fmt.Println("Sum of odd numbers: ", oddtotal)
}

func sum(num ...int) int {
	t := 0
	for _, v := range num {
		t += v
	}
	return t
}

func even(f func(xi ...int) int, x ...int) int {
	var y []int
	for _, v := range x {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

func odd(f func(xi ...int) int, x ...int) int {
	var y []int
	for _, v := range x {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
