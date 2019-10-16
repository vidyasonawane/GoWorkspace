package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) //max is not the function anymore
	//fmt.Println(max(8)) //This will give the error as max is a variable now.
}

//Don't do this, bad coding practice to shadoe variables.