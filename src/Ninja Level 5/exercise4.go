//anonymous structs
package main

import "fmt"

func main() {

	p := struct {
		name  string
		age   int
		happy bool
	}{
		name:  "vidya",
		age:   22,
		happy: true,
	}

	fmt.Println(p)

}
