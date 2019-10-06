package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	p1 := person{
		first: "vidya",
		last:  "sonawane",
		favFlavors: []string{
			"vanilla",
			"chocolate",
		},
	}
	p2 := person{
		first: "kunal",
		last:  "sonawane2",
		favFlavors: []string{
			"strawberry",
			"butterscotch",
		},
	}

	m:= map[string]person {
		p1.last: p1,
		p2.last:p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, f := range v.favFlavors {
			fmt.Println(i,f)
		}
	}
}
