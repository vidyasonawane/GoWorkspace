package main

import "fmt"

func main() {
	//array with default value
	var myarr1 [5]int
	fmt.Println(myarr1)
	myarr1[3] = 12
	fmt.Println(myarr1)

	//using composite literals
	myarr2 := [4]int{1, 2, 3, 4}
	//var myarr2 = [4]int {1,2,3,4}
	fmt.Println(myarr2)

	//... instead of array size
	myarr3 := [...]int{2, 4, 6, 8, 10, 1, 2, 14, 16}
	fmt.Println(myarr3)
	fmt.Println(len(myarr3))

	//In arrays, the [size] should not be empty
}
