package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Inside speak", p.first, p.last, p.age)
}
func main() {
	p1 := person{
		first: "vidya",
		last:  "sonawane",
		age:   22,
	}

	p1.speak()
}
