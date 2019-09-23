package main

import "fmt"

func main() {
	for i := 33; i < 127; i++ {
		fmt.Printf("%d\t%x\t%#U\n", i, i, i)
	}
}

// %#U = to print ascii
