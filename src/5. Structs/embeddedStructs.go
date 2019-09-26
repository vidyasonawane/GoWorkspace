package main

import "fmt"

type person struct {
	first string
	dob   int
}

type sagent struct {
	person
	ltk bool
}

func main() {

	sa1 := sagent{
		person: person{
			first: "vidya",
			dob:   1997,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.first)
	//or
	fmt.Println(sa1.person.first)

	fmt.Println(sa1.dob)
	//or
	fmt.Println(sa1.person.dob)

	fmt.Println(sa1.ltk)

}
