package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("not printed")
	case true:
		fmt.Println("Its true")
		// fallthrough
	case 2 == 2:
		fmt.Println("Its equal")
		// fallthrough
	default:
		fmt.Println("In default")

	}

	switch "Bond" {
	case "James":
		fmt.Println("James")
	case "Bond":
		fmt.Println("Its Bond")
	default:
		fmt.Println("No one")

	}

}
