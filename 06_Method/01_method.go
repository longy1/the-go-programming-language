package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (*Point) Foo() {
	fmt.Printf("Foo in Point")
}

type Path []Point

func (path Path) Distance() float64 {
	var sum float64
	for i, p := range path {
		if i > 0 {
			sum += p.Distance(path[i-1])
		}
	}
	return sum
}

func main() {
	p1 := Point{3, 4}
	p2 := Point{6, 2}
	fmt.Println(p1.Distance(p2))
	fmt.Println((&p1).Distance(p2))

	path := Path{{1, 2}, {4, 5}, {7, 2}, {3, 4}}
	fmt.Println(path.Distance())

}
