package main

import "fmt"

func main() {
	x := make([]string, 5, 50)

	x[0] = "gulab jamun"
	x[1] = "Ras malai"
	x[2] = "Kaju katli"
	x[3] = "Rassogulla"
	x[4] = "Aamras"
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
