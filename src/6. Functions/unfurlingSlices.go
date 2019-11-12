package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := foo(xi...)
	fmt.Println("Sum in main : ", sum)
}
func foo(x ...int) int { //unlimited number of int
	fmt.Println(x)
	fmt.Printf("%T\n", x) //slice of int

	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	fmt.Println("Sum is :", sum)
	return sum
}
