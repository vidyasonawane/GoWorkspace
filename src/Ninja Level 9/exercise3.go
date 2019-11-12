package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()

		fmt.Println("Num of gs: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("....", counter)
	fmt.Println("Num of gs at end: ", runtime.NumGoroutine())

	fmt.Println("About to exit")
}
