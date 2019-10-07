package main

import "fmt"
import "math"

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(ss shape) {
	fmt.Println("Area of ", ss, " is : ", ss.area())
}

func main() {
	c := circle{
		radius: 12.345,
	}
	s := square{
		length: 2.5,
	}

	info(c)
	info(s)
}
