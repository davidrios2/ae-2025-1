package _1_OCP

import (
	"fmt"
	"math"
)

// Shape interface defines a contract for any shape that can calculate its area.
type Shape interface {
	Area() float64
}

//There is already another struck with this name, so we can not have multiple strucks with the same name
// Rectangle struct.
/*type Rectangle struct {
	Width  float64
	Height float64
}*/
// Circle struct
/*type Circle struct {
	Radius float64
}*/

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// CalculateArea
/*type AreaCalculator struct{}*/

func (c *AreaCalculator) CalculateArea(shape Shape) float64 {
	return shape.Area()
}

func OcpC() {
	rectangle := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}
	triangle := Triangle{Base: 6, Height: 4}

	calculator := AreaCalculator{}
	rectangleArea := calculator.CalculateArea(rectangle)
	circleArea := calculator.CalculateArea(circle)
	triangleArea := calculator.CalculateArea(triangle)

	fmt.Println("Rectangle Area:", rectangleArea)
	fmt.Println("Circle Area:", circleArea)
	fmt.Println("Triangle Area:", triangleArea)
}
