package main

import "fmt"

func main() {
	s := "Hello Vidya"
	fmt.Println(s)
	fmt.Printf("Using format printing: %s\n", s)

	t := `Hey "Vidya"
	How you Doing?`
	fmt.Println("Using back tick : \n", t)

	bs := []byte(s) //Type conversion
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for k, val := range s {
		fmt.Printf("%v %v\n", k, val)
	}

}
