package main

import "fmt"

// const a = 40
// const b = 12.34
// const c = "vidya"
// const d int //error, const dont have zero values. it should be defined always.

// const (
// 	a = 40
// 	b = 12.34
// 	c = "vidya"
// )

const (
	a int     = 40
	b float64 = 12.34
	c string  = "vidya"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	// const d := 234 //error, we cannot use short declaration operator with const.
	const d = 234
	fmt.Println(d)
}
