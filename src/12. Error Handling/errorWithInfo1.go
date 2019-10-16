package main

import (
	"errors"
	"fmt"
	"log"
)

var myerr = errors.New("This error is created by me")

func main() {
	val, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(val, err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, myerr
	}
	return 22, nil
}
