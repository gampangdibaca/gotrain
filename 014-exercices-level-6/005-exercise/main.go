package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	r float64
}

func (s square) area() {
	fmt.Println("The total area of square is", s.length*s.width)
}

func (c circle) area() {
	if int(c.r)%7 == 0 {
		fmt.Println("The total area of circle is", 22*(c.r*c.r/7))
	} else {
		fmt.Println("The total area of circle is", math.Pi*c.r*c.r)
	}
}

type shape interface {
	area()
}

func info(s shape) {
	switch s.(type) {
	case square:
		s.(square).area()
	case circle:
		s.(circle).area()
	}
}

func main() {
	s := square{
		length: 5,
		width:  8,
	}
	c := circle{
		r: 7,
	}
	info(s)
	info(c)
}
