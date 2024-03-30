package main

import (
	"fmt"
	"math"
)

// interfaces are implemented implicitly by the structs
type Shape interface {
	area() float64
}

type Rectangle struct  {
	x1, x2, y1, y2 float64
} 

type Circle struct  {
	x float64
	y float64
	r float64
}

func main() {
  var rectangle Rectangle = Rectangle{
		x1: 2,
		x2: 5,
		y1: 10,
		y2: 11,
	}

	var circle = Circle{
		x: 0,
		y: 0,
		r: 5,
	}

  fmt.Println(rectangleArea(&rectangle))
  fmt.Println(rectangle.area())
  fmt.Println(circleArea(&circle))
  fmt.Printf("Total area %f\n", totalArea(&circle, &rectangle))
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}

	return area
}
 
func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1

  return math.Sqrt(a*a + b*b)
}

func rectangleArea(rectangle *Rectangle) float64 {
  l := distance(rectangle.x1, rectangle.y1, rectangle.x1, rectangle.y2)
  w := distance(rectangle.x1, rectangle.y1, rectangle.x2, rectangle.y1)
  return l * w
}

func circleArea(circle *Circle) float64 {
  return math.Pi * circle.r * circle.r
}

// write a method to the struct directly, this implicitly will have the &r passed to it when called through dot method - i.e. rectangle.area() 
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}