package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First  string
	Last   string
	Saying []string
}

func main() {
	p1 := person{
		First:  "James",
		Last:   "Bond",
		Saying: []string{"Shaken, not stirred", "Any last wish"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error message is: ", err)
	}
	fmt.Println(string(bs))
}
