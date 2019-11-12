package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 22
	c <- 23

	fmt.Println(<-c)
	fmt.Println(<-c)
}

//fatal error: all goroutines are asleep - deadlock!
//channel can only store one value.
