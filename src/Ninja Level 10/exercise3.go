package main

import "fmt"
//Range Over channels and close them

func main() {

	c := make(chan int)

	//we are using goroutine because sending anf receiving on channel should happen at same time.
	go func(){ 
			for i := 0;i<10;i++{
			c <- i
		}
		close(c) //if it is not closed then gives fatal error:deadloack
	}()
	
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to exit")

}