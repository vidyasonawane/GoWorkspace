package main

import "fmt"

func main() {
	xi := []int{3, 4, 5, 6,7,8}
	total := foo(xi...)
	fmt.Println(total)
	
	total2 := bar(xi)
	fmt.Println(total2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(y []int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}
