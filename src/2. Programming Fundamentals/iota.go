package main

import "fmt"

// const (
// 	a = iota
// 	b
// 	c
// )
// op = 0 1 2

// const (
// 	a = 10
// 	b = 22
// 	c = iota
// 	d
// )
// op = 10 22 2 3

const (
	a = "vidya"
	b = "kunal"
	c = iota
	d
)

const x = iota

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
