package main

import "fmt"

func main() {
	x := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 10}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
