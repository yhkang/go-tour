package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

//the receiver (p) behaves exactly as if it were an argument to the method
func (p *Point) ScaleByPointer(factor float64) {
    p.X *= factor
    p.Y *= factor
}

func main() {
	point := Point{1,2}
	fmt.Println(point)

	point.ScaleBy(2)
	fmt.Println(point)

	point.ScaleByPointer(2)
	fmt.Println(point)
}