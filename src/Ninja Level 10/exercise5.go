package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 22
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c) //try running code by commenting this line.

	v, ok = <-c
	fmt.Println(v, ok)
}
