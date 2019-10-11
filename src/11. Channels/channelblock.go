package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 22
	fmt.Println(<-c)
}

//Output:fatal error: all goroutines are asleep - deadlock!
//The sending and receiving should occur are one time, for that we require 2 goroutines, but here is only one which is main.