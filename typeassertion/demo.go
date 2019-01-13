package main

import (
	"fmt"
	"math"
)

func main() {
	var shapeList = new([]Shape)
	*shapeList = append(*shapeList, Circle{1})
	*shapeList = append(*shapeList, Circle{5})
	*shapeList = append(*shapeList, Rectangle{4, 5})

	calcArea(shapeList)
}

type Shape interface {
	Area() float32
}

type Circle struct {
	r float32
}

type Rectangle struct {
	width  float32
	length float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func calcArea(list *[]Shape) {
	for _, shape := range *list {
		_, isCircle := shape.(Circle)
		if isCircle {
			fmt.Println("A circle")
		}
		_, isRecta := shape.(Rectangle)
		if isRecta {
			fmt.Println("A rectangle")
		}
		fmt.Println(shape.Area())
	}
}
