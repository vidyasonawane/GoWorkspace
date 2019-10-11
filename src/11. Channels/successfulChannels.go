package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 22
	}()

	fmt.Println(<-c)
}
