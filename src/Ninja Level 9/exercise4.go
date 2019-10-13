package main

//mutex to avoid race condition
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()

		fmt.Println("Num of gs: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("....", counter)
	fmt.Println("Num of gs at end: ", runtime.NumGoroutine())

	fmt.Println("About to exit")
}
