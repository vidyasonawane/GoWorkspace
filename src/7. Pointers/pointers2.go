package main

import "fmt"
//using pointers
func main() {
	x := 0
	fmt.Println("x before: ", x)
	fmt.Println("x before: ", &x)
	foo(&x)
	fmt.Println("x after: ", x)
	fmt.Println("x after: ", &x)

}

func foo(y *int) {
	fmt.Println("y before: ", *y)
	fmt.Println("y before: ", y)
	*y = 22
	fmt.Println("y after: ", *y)
	fmt.Println("y after: ", y)
}
