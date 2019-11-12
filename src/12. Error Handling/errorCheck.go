package main

import "fmt"

func main() {
	n, err := fmt.Println("Vidya")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n) //5 of vidya + 1 of new line = 6
}
