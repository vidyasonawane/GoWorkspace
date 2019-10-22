package main

import "fmt"

type Contact struct {
	name string
	age int
}

func switchOnType(x interface{}) {
	switch x.(type) { //this is asserting
	case int:
		fmt.Println("Int",x)
	case string:
		fmt.Println("String",x)
	case Contact:
		fmt.Println("Struct contact",x)
	default:
		fmt.Println("default")
	}
}

func main() {
	switchOnType(22)
	switchOnType("vidya")
	switchOnType(Contact{"vidya",22})
}
