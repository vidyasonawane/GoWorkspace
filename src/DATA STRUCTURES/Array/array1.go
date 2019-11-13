package main

import "fmt"

func main() {

	var x [256]byte

	fmt.Println(len(x))

	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}

	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}

}
