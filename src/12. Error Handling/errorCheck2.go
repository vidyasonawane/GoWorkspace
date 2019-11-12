package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//Writing into file
func main() {
	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f, r)
}
