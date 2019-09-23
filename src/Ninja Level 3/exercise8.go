package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("Prints")
	case false:
		fmt.Println("Not printing")
	default:
		fmt.Println("Default")
	}
}
