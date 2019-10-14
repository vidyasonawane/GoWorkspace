package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string
	//Scan returns the number of items successfully scanned and error
	fmt.Println("Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Fav food: ")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Fav Sport: ")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nEntered values are :")
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)

}
