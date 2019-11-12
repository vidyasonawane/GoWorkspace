package main

import "fmt"

func main() {
	myarr := [5]int{1, 2, 3, 4, 5}

	for i, v := range myarr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", myarr)

}
