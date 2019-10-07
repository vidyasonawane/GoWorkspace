package main

import "fmt"

func main() {
	fmt.Println("In main")
	func() {
		fmt.Println("I am in anonymous function")
	}()

	func(x int) {
		fmt.Println(x)
	}(28)
}
