package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 22

	fmt.Println(<-c)
}

//No need of another goroutine because the channel is able to store one value