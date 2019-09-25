package main

import "fmt"

func main() {

	ms := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"miss", "MoneyPenny", "Hellooooo, James."}}

	fmt.Println(ms)

	for i, ms1 := range ms {
		fmt.Println("Record: ", i)
		for j, val := range ms1 {
			fmt.Println(j, val)
		}
	}
}
