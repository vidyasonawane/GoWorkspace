//In addition to main goroutine, launch two additional goroutines
//each additional goroutine should print something out
//use waitgroups to make sure each goroutinne finishes before program exists

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU at beginning: ", runtime.NumCPU())
	fmt.Println("Goroutines at beginning: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("In first goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("In second goroutine")
		wg.Done()
	}()

	fmt.Println("CPU at mid: ", runtime.NumCPU())
	fmt.Println("Goroutines at mid: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("About to exit")
	fmt.Println("CPU at end: ", runtime.NumCPU())
	fmt.Println("Goroutines at end: ", runtime.NumGoroutine())

}
