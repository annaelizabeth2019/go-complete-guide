package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := &square{
		sideLength: 5,
	}
	t := &triangle{
		base:   64,
		height: 5,
	}
	printArea(sq)
	printArea(t)
}

func (s *square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}
func (t *triangle) getArea() float64 {
	area := t.base * t.height
	return area
}

func printArea(s shape) {
	var area float64
	area = s.getArea()
	fmt.Println(area)
}
