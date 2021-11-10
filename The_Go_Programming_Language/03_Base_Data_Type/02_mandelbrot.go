// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/mandelbrot", mandelbrotHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	xmin, ymin, xmax, ymax := -2.0, -2.0, +2.0, +2.0
	width, height := 1024.0, 1024.0
	xForm := r.FormValue("x")
	yForm := r.FormValue("y")
	zoomForm := r.FormValue("zoom")
	if xForm != "" {
		temp, _ := strconv.Atoi(xForm)
		xmax = float64(temp)
		xmin = -xmax
	}
	if yForm != "" {
		temp, _ := strconv.Atoi(yForm)
		ymax = float64(temp)
		ymin = -ymax
	}
	if zoomForm != "" {
		zoom, _ := strconv.ParseFloat(zoomForm, 64)
		width = zoom * width
		height = zoom * height
	}

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < int(width); px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, colorfulMandelbrot(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

func colorfulMandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: 35, G: 118, B: 183, A: 255 - contrast*n}
		}
	}
	return color.RGBA{R: 20, G: 74, B: 116, A: 200}
}
