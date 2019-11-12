package main

import "fmt"

func main() {
	x := "Bond"
	if true {
		fmt.Println("True")
	}
	if false {
		fmt.Println("False")
	}
	if x == "Bond" {
		fmt.Println(x, ", James Bond")
	}
}
