package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func bar(h human) {
	switch h.(type) { //Assertion
	case person:
		fmt.Println("I was passed here", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed here", h.(secretAgent).first)

	}
	fmt.Println("I was passed into bar", h)
}

type human interface {
	speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "vidya",
			last:  "sonawane",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "kunal",
			last:  "patil",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(p1)
	sa1.speak()
	sa2.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	//conversion

}

func (r secretAgent) speak() {
	fmt.Println("Hey", r.first, r.last)
}

func (p person) speak() {
	fmt.Println("Hey", p.first, p.last)
}
