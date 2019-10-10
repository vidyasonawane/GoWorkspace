package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 6, 1, 3, 8, 0, 3}
	xs := []string{"vids", "james", "vira", "kunal"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("----------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
