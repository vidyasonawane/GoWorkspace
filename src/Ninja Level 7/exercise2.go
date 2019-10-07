package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "James",
		age:  22,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}

func changeMe(p *person) {
	p.name = "vidya"
	//or
	//(*p).name = "vids"

}
