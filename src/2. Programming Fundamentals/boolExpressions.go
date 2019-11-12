package main

import "fmt"

func main() {
	if true || false {
		fmt.Println("OR operator works!")
	}

	if true && false {
		fmt.Println("AND operator works!")
	}


}