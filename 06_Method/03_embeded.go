package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point2 struct {
	X, Y float64
}

func (p *Point2) Distance(q Point2) float64 {
	fmt.Println("Point.Distance")
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type ColorfulPoint struct {
	Point2
	color color.RGBA
}

func (cp *ColorfulPoint) Distance(q ColorfulPoint) float64 {
	fmt.Println("ColorfulPoint.Distance")
	return math.Hypot(cp.X-q.X, cp.Y-q.Y)
}

func main() {
	p := ColorfulPoint{Point2{2, 3}, color.RGBA{}}
	q := ColorfulPoint{Point2{4, 5}, color.RGBA{}}
	fmt.Println(p.Distance(q))
}
