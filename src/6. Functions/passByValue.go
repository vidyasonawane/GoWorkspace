package main

import "fmt"

func main() {
	//Everything in Go is pass by value
	age := 22

	fmt.Println("Before ChangeMe(): ", age)
	ChangeMe(age) //here we are passing age
	fmt.Println("After ChangeMe(): ", age)

	fmt.Println("Before ChangeMeUsingAddress(): ", &age, age)
	ChangeMeUsingAddress(&age) //here we are passing the address as a value where the age is stored.
	//We should pass the address only for non reference types like int, float64.
	// the slices, maps and channels are already reference types hence no need to pass value of address.
	fmt.Println("After ChangeMeUsingAddress(): ", &age, age)
	
}

func ChangeMe(a int) {
	a = 25
}

func ChangeMeUsingAddress(z *int) {
	fmt.Println("inside ChangeMeUsingAddress")
	fmt.Println(z)
	fmt.Println(*z)
	*z = 44
	fmt.Println(z)
	fmt.Println(*z)
	fmt.Println("Leaving ChangeMeUsingAddress()")
}
