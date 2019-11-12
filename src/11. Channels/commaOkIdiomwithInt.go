package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("About to exit")
}

func send(even, odd, quit chan int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Received Even Value: ", v)
		case v := <-odd:
			fmt.Println("Received Odd value : ", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("From comma ok idiom  : ", i, ok)
				return
			} else {
				fmt.Println("From comma ok idiom: ", i)
			}
		}
	}
}
