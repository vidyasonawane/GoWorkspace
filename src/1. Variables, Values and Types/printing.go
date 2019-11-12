package main

import "fmt"

func main() {
	x := 100
	fmt.Print(x)
	fmt.Println(x)
	fmt.Printf("%v\n", x)  //original value
	fmt.Printf("%b\n", x)  //binary
	fmt.Printf("%x\n", x)  //hexadecimal
	fmt.Printf("%#x\n", x) //hexadecimal starting with 0x
	fmt.Printf("%T\n", x)  //gives the type

	y := fmt.Sprintln(x)
	fmt.Println(y)
	fmt.Printf("%T\n\n", y)

	z := fmt.Sprintf("%v", 10)
	fmt.Println(z)
	fmt.Printf("%T", z)

}
