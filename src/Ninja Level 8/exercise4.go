package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 7, 3, 1, 6, 0, 3, 66, 12, 45, 89, 90}
	fmt.Println("Before Sorting: ", s)
	sort.Ints(s)
	fmt.Println("After Sorting:  ", s)

	str := []string{"vidya", "vids", "chakuli", "gabadi", "rancho", "doblya", "dambis", "mamma"}
	fmt.Println("Before Sorting: ", str)
	sort.Strings(str)
	fmt.Println("After Sorting:  ", str)
}
