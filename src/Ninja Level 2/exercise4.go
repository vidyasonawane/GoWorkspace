package main

import "fmt"

func main() {
	var a int = 47
	fmt.Printf("%b\t%d\t%x\n", a, a, a)
	b := a << 1
	fmt.Printf("%b\t%d\t%x\n", b, b, b)
}
