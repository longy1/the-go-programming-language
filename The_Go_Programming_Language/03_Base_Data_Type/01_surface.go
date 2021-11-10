// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-max, ..., max)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes, = 30°
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/surface", SurfaceHandler)
	http.HandleFunc("/eggbox", EggBoxHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func EggBoxHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	DrawEggBox(w)
}

func DrawEggBox(w io.Writer) {
	if _, err := fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height); err != nil {
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := cornerEggBox(i+1, j)
			bx, by := cornerEggBox(i, j)
			cx, cy := cornerEggBox(i, j+1)
			dx, dy := cornerEggBox(i+1, j+1)
			if _, err := fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy); err != nil {
			}
		}
	}
	if _, err := fmt.Fprintf(w, "</svg>"); err != nil {
	}
}

func cornerEggBox(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute height z
	z := zEggBox(x, y)

	// compute 2-D canvas point
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale + 100
	return sx, sy
}

func zEggBox(x, y float64) float64 {
	x = math.Mod(math.Abs(x), 0.5) - 0.25
	y = math.Mod(math.Abs(y), 0.5) - 0.25

	r := math.Exp(math.Hypot(x, y))
	return r
}

func SurfaceHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	DrawSurface(w)
}

func DrawSurface(w io.Writer) {
	if _, err := fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height); err != nil {
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := cornerSurface(i+1, j)
			bx, by := cornerSurface(i, j)
			cx, cy := cornerSurface(i, j+1)
			dx, dy := cornerSurface(i+1, j+1)
			if _, err := fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy); err != nil {
			}
		}
	}
	if _, err := fmt.Fprintf(w, "</svg>"); err != nil {
	}
}

func cornerSurface(i, j int) (float64, float64) {
	// find point(x, y) for cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute height z
	z := zSurface(x, y)

	// compute 2-D canvas point
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func zSurface(x, y float64) float64 {
	r := math.Hypot(x, y) // == sqrt(x^2 + y^2)
	z := math.Sin(r) / r
	if z > math.MaxFloat64 {
		return math.MaxFloat64
	}
	return math.Sin(r) / r
}
