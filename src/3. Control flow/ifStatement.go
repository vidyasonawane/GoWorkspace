package main

import "fmt"

func main() {
	if true {
		fmt.Println("True")
	}
	if false {
		fmt.Println("False, never prints")
	}
	if 2 == 2 {
		fmt.Println("Equal")
	}
	if !(2 == 2) {
		fmt.Println("Not equal, never prints")
	}

	//to limit the scope of a for if else only
	if a := 9; a > 9 {
		fmt.Println("a is greater than 9")
	} else if a == 9 {
		fmt.Println("a is equal to 9")
	} else {
		fmt.Println("a is less than 9")
	}

}
