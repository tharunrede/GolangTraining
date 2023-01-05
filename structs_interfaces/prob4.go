package main

import "fmt"

type shape interface {
	area() float64
	//peri() float64
}

func info(i shape) float64 {
	return i.area()
}

type square struct {
	length int
	width  int
}

type circle struct {
	radius int
}

func (sq square) area() float64 {

	return float64(sq.length) * float64(sq.width)

}

func (sq square) peri() float64 {

	return float64(2 * (sq.length + sq.width))

}

func (cir circle) area() float64 {

	return 3.14 * float64(cir.radius) * float64(cir.radius)

}

func main() {

	square1 := square{4, 5}
	circle1 := circle{5}
	fmt.Println("area of square is : ", square1.area())
	fmt.Println("area of square is : ", circle1.area())

	//interfacevar := []shape{square1, circle1}

	fmt.Println("Square area using interface", info(square1))
	fmt.Println("circle area using interface", info(circle1))

	//fmt.Println(interfacevar)

}
