package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("nofile.txt")
	if err != nil {
		fmt.Println("Error msg using printing: ", err)
		log.Println("Error msg using logging: ", err) //with timestamp
		log.Fatalln(err)
		log.Panicln(err)
		panic(err)
	}
}
