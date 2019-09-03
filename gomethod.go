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


type ColoredPoint struct {
    *Point
    color string
}

func (p *ColoredPoint) ScaleBy(factor float64) {
    p.Point.ScaleByPointer(factor)
}

func main() {
	point := Point{1,2}
	fmt.Println(point)

	point.ScaleBy(2)
	fmt.Println(point)

	point.ScaleByPointer(2)
	fmt.Println(point)

	p := ColoredPoint{&Point{1, 1}, "red"}
	q := ColoredPoint{&Point{5, 4}, "blue"}
	fmt.Println(*q.Point) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}