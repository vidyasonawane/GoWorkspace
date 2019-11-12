package main

import "fmt"

func main() {
	//Closures are used to enclose the scope of variable for certain area.
	//variables declared in outer scope can be accessible in inner scope.
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int //No need to pass it to function, its scope is inside the function.
	return func() int {
		x++
		return x
	}
}
