package main

import "fmt"

func main() {
	a := myfunc()
	fmt.Println(a())
}

func myfunc() func() int {
	return func() int {
		return 10
	}
}
