package _1_OCP

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type AreaCalculator struct{}

func (c *AreaCalculator) CalculateRectangleArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func (c *AreaCalculator) CalculateCircleArea(circle Circle) float64 {
	return 3.14159 * circle.Radius * circle.Radius
}

func OcpI() {
	rectangle := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}

	calculator := AreaCalculator{}
	rectangleArea := calculator.CalculateRectangleArea(rectangle)
	circleArea := calculator.CalculateCircleArea(circle)

	fmt.Println("Rectangle Area:", rectangleArea)
	fmt.Println("Circle Area:", circleArea)
}
