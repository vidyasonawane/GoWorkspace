package main

import "fmt"

var z uint8 // also called as byte
var t int8
var s int32 //also called as rune

func main() {
	a := 12
	fmt.Printf("%T\n", a)

	b := 12.34
	fmt.Printf("%T\n", b)
	fmt.Printf("%2.4f\n", b) //2 width and 4 precision

	// z = -127 //error because it is unsigned
	// z = 256 // error bcoz range is 0 to 255
	z = 29
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	t = -24 // range is -128 to 127
	fmt.Println(t)

}
