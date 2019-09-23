package main

import "fmt"

func main() {
	fmt.Printf("\ntrue && true: %v", true && true)
	fmt.Printf("\ntrue && false: %v", true && false)
	fmt.Printf("\ntrue || true: %v", true || true)
	fmt.Printf("\ntrue || false: %v", true || false)
	fmt.Printf("\n!true: %v", !true)
	fmt.Printf("\n!false: %v", !false)
}
