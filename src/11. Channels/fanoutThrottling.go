package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("About to exit")

}

func populate(c1 chan int) {
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const numGoroutines = 10 //only 1o goroutines at a time
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeconsuming(v2)
				}(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(c2)
}

func timeconsuming(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
