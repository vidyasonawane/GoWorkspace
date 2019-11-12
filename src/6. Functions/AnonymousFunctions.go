package main

import "fmt"

func main() {

	foo()

	func() {
		fmt.Println("I am in anonymous function")
	}()

	func(x int) {
		fmt.Println("I am in anonymous function with parameters", x)
	}(22)

}

func foo() {
	fmt.Println("I am in foo")
	func() { //self executing function.
		fmt.Println("I am in anonymous function in foo")
	}()
}
