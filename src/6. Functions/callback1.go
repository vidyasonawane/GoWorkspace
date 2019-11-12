package main

import "fmt"

func visit(numbers []int, callbackFunc func(int)) {
	for _, n := range numbers {
		callbackFunc(n)
	}
}
func main() {
	//callback is passing an function as an argument
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}
