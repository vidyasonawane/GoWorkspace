package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("when %d is divided by 4, remainder is %d\n", i, i%4)
	}
}
