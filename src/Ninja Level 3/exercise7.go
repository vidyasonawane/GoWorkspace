package main

import "fmt"

func main() {
	x := "Vidya"
	if x == "James" {
		fmt.Println("James")
	} else if x == "Bond" {
		fmt.Println("Bond")
	} else {
		fmt.Println("No bond ;)", x)
	}
}
