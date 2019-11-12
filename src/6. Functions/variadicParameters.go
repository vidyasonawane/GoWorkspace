package main

import "fmt"
//parameters = while function definition
func main() {
	sum := foo(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Sum in main : ", sum)

	//variadic parameters means zero or more
	sum2 := foo()
	fmt.Println("Sum in main : ", sum2)
}
func foo(x ...int) int { //unlimited number of int (zero or more)
	fmt.Println(x)
	fmt.Printf("%T\n", x) //slice of int

	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	fmt.Println("Sum is :", sum)
	return sum
}
