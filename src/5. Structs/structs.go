package main

import "fmt"

type person struct {
	first string
	dob   int
}

func main() {
	p1 := person{
		first: "vidya",
		dob:   1997,
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.dob)

}
