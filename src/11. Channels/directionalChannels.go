package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Printf("bidirectional: %T\t", c)
	fmt.Printf("\nReceive only: %T\t", cr)
	fmt.Printf("\nSend only: %T\t", cs)
}
