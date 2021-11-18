package main

import "fmt"

type Point3 struct {
	X, Y float64
}

func (p Point3) Add(q Point3) Point3 { return Point3{p.X + q.X, p.Y + q.Y} }
func (p Point3) Sub(q Point3) Point3 { return Point3{p.X - q.X, p.Y - q.Y} }

type Path2 []Point3

func (path Path2) TranslateBy(offset Point3, add bool) {
	var op func(p, q Point3) Point3
	if add {
		op = Point3.Add
	} else {
		op = Point3.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	path := Path2{{2, 3}, {2, 4}, {3, 3}}
	fmt.Println(path)
	path.TranslateBy(Point3{1, 2}, true)
	fmt.Println(path)
	path.TranslateBy(Point3{2, 1}, false)
	fmt.Println(path)
}
