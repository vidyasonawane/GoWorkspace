package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("In foo")
}

func bar() {
	fmt.Println("In bar")
}

// without defer:
// In foo
// In bar

// with defer:
// In bar
// In foo
