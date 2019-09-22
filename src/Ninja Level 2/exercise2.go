package main

import "fmt"

func main() {
	a := (11 == 12)
	b := (11 <= 12)
	c := (11 >= 12)
	d := (11 != 12)
	e := (11 < 12)
	f := (11 > 12)

	fmt.Println(a, b, c, d, e, f)
}
