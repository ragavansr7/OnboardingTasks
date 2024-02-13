package main

import (
	"fmt"
	"math"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// func Area(rectangle Rectangle) float64{
// 	return rectangle.Height*rectangle.Width
// }

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func main() {
	fmt.Println(Hello("ragavan"))
	fmt.Println(Perimeter(Rectangle{2.0, 2.0}))
}
