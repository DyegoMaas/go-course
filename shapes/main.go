package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangleRectangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func (t triangleRectangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func main() {
	s := square{
		side: 10.5,
	}
	printShape(s)

	t := triangleRectangle{
		base:   10.5,
		height: 5.8,
	}
	printShape(t)
}

func printShape(s shape) {
	fmt.Println("Shape area is", s.getArea())
}
