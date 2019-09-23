package main

import "fmt"

func main() {
	//3 ways to print for loop

	fmt.Println("way 1 : ")
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("way 2 : ")
	j := 1
	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("way 3 : ")
	k := 1
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	fmt.Println("nested loops : ")

	for x := 0; x < 2; x++ {
		for y := 0; y < 4; y++ {
			fmt.Println(x, y)
		}
	}

}
