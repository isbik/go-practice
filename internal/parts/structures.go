package parts

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	w, h float64
}

func (r *Rectangle) area() float64 {
	return r.h * r.w
}

func (r *Rectangle) perimeter() float64 {
	return (r.h + r.w) * 2
}

func Structures() {
	r := Rectangle{100, 500}
	c := Circle{500}

	fmt.Println(totalArea(&r, &c))
	fmt.Println("--------------")
	fmt.Println(totalPerimeter(&r, &c))
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.perimeter()
	}

	return area
}
