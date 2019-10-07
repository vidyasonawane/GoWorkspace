package main

import "fmt"

func main() {

	defer foo()
	fmt.Println("In main")
	bar()
	fmt.Println("At the end of main")
}

func foo() {
	defer func() {
		fmt.Println("In anonymous function inside foo")
	}()
	fmt.Println("in foo")
}

func bar() {
	fmt.Println("in bar")
}
