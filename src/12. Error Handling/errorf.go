package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-12.678)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("This is is created by me %v", f)
	}
	return 22, nil
}
