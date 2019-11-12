package main

import "fmt"
import "math"

//will not work
//pointer receiver and non pointer value
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {
	c := circle{5}
	info(c)
}
