package main

import "fmt"

//Range Over channels and close them using function and retuen

func main() {

	c := gen()
	printchannel(c)
	fmt.Println("About to exit")

}

//return receive only channel
func gen() <-chan int {
	c := make(chan int)

	//we are using goroutine because sending anf receiving on channel should happen at same time.
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) //if it is not closed then gives fatal error:deadloack
	}()
	return c
}

func printchannel(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}

}
