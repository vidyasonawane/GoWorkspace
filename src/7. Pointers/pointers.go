package main

import "fmt"

func main() {
	a := 22
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	//var b int = &a //gives error
	var b *int = &a
	//or
	//b := &a
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Println(*b)
	*b = 99
	fmt.Println(a)

}
