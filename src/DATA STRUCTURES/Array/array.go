package main

import "fmt"

func main() {
	//Array is a numbered sequence of elements of same type
	// the number of elements is length which is the part of an array.
	//Not dynamic: does not change in size.
	var x [58]string
	fmt.Println(x)

	fmt.Println(x[43]) //gives 44th item because of zero based index.

	for i := 65; i<=122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[43])

}