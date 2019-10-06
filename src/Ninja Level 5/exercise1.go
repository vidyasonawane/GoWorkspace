package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	p1 := person {
		first: "vidya",
		last:  "sonawane",
		favFlavors : []string {
			"vanilla",
			"chocolate",
		},
	}
	p2 := person {
		first: "kunal",
		last:  "sonawane",
		favFlavors: []string {
			"strawberry",
			"butterscotch",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for i,v := range p1.favFlavors{
		fmt.Println(i,v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)

	for i,v := range p2.favFlavors{
		fmt.Println(i,v)
	}

}
