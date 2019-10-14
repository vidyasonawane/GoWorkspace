package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("nofile.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit(1) is called, deferred functions don't run")
}
