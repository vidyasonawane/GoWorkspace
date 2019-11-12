//Everything in go is pass by value.
//We don't pass value of address in case of reference types because they are already referencing to some address.

//Reference Types: Slice, map, channels
//Non refernce types; int, float64, bool, string.

package main

import "fmt"

func main() {
	// Slice
	s := make([]string, 1, 25) //type length capacity
	fmt.Println(s)
	ChangeMe(s)
	fmt.Println(s)

	//Map
	m := make(map[string]int)
	ChangeMap(m)
	fmt.Println(m["vidya"])
}

func ChangeMe(z []string) {
	z[0] = "Vidya"
	fmt.Println(z) //22
}

func ChangeMap(x map[string]int) {
	x["vidya"] = 22
}
