package main

import "fmt"

func main() {
	//To avoid deadloack
	//1. use goroutines
	c := make(chan int)
	go func() {
		c <- 22
	}()

	fmt.Println(<-c)

	//OR 
	//2. buffered channels (not an ideal one though :()
	c2 := make(chan int,1)
	c2 <- 44
	fmt.Println(<-c2)


}
