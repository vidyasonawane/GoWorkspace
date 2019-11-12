package main

import "fmt"

func main() {
	s1 := foo() //returning a string
	fmt.Println(s1)

	x := bar() //returning a function

	fmt.Printf("%T\n", x)

	fmt.Println(x)
	i := x()
	fmt.Println(i)

	fmt.Println(greet("vidya","kunal"))
}

//returning a value
func foo() string {
	return "Vidya"
}

//returning a function
func bar() func() int {
	return func() int {
		return 2019
	}
}

//returning multiple values
func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname,lname), fmt.Sprint(lname,fname)
}
