package main

import "fmt"

//Arguments: while calling the function

func main() {
 data := []float64{2,4,6,8,10}
 n := average(data...)
 fmt.Println(n)
}

func average(sf...float64) float64 {
	var total float64
	for _,v := range sf {
		total += v
	}

	return total/float64(len(sf))
}