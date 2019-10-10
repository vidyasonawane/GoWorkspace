package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "vidya",
		Last:  "sonawane",
		Age:   22,
	}
	p2 := person{
		First: "kunal",
		Last:  "patil",
		Age:   19,
	}

	people := []person{p1, p2} //slice of person

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
