package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func main() {
	fmt.Println("Main start")
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	fmt.Println("Inside sqrt")
	if f < 0 {
		nme := fmt.Errorf("norget mathredux: square root of negative number")
		return 0, norgateMathError{"50.483 N", "99.234 W", nme}
	}
	return 42, nil
}

func (n norgateMathError) Error() string {
	fmt.Println("Inside Error")
	return fmt.Sprintf("A norgate math erroe occured: %v %v %v", n.lat, n.long, n.err)
}
