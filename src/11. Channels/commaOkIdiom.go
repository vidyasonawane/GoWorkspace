package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	//channel is closed
	v, ok = <-c
	fmt.Println(v, ok)
}
