package main

import "fmt"

func main() {
	foo()
	bar("vidya")
	
	s1 := woo("Kunal")
	fmt.Println(s1)

	x, y := myfunc("vidya", 22)
	fmt.Println(x)
	fmt.Println(y)
}

//basic func
func foo() {
	fmt.Println("In foo")
}

//function with parameters
//Everything in go is pass by value.
func bar(s string) {
	fmt.Println("In bar", s)
}

//function with parameters and return type
func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

//function with multiple return and multiple parameters

func myfunc(name string, age int) (string, bool) {
	str := fmt.Sprint("Hello ", name, ", your age is ", age)
	b := true
	return str, b
}
