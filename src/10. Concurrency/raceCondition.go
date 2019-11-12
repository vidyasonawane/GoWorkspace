package main

//Run this using terminal and notepad
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

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
			wg.Done()
		}()
		fmt.Println("Goroutines = ", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

//To detect the race condition, run the following command
//go run -race racecondition.go
