package main

import "fmt"

func main() {
	//cs := make(chan<- int) //send only channel
	cs := make(chan int) //this will work
	go func() {
		cs <- 22
	}()

	fmt.Println(<-cs) //gives error in 1st case as it is receiving from send only channel
	fmt.Printf("cs\t%T\n", cs)

}
