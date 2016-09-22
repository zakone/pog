//server of mandelbrot function
//http://localhost:8000/?x=0.5&y=0.5&scale=2

package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "image/png")
		draw(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Params struct {
	delta  float64
	x      float64
	y      float64
	width  int
	height int
	scale  int
}

var param Params

func draw(w io.Writer, r *http.Request) {

	param.delta = 2.0
	param.x = 0
	param.y = 0
	param.scale = 1
	param.width = 1024
	param.height = 1024

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "x" {
			param.x, _ = strconv.ParseFloat(v[0], 64)
			log.Printf("center x is %f", param.x)
		} else if k == "y" {
			param.y, _ = strconv.ParseFloat(v[0], 64)
			log.Printf("center y is %f", param.y)
		} else if k == "scale" {
			param.scale, _ = strconv.Atoi(v[0])
			log.Printf("scale to %d", param.scale)
		}
	}
	width := param.width * param.scale
	height := param.height * param.scale
	xmin := param.x - param.delta
	xmax := param.x + param.delta
	ymin := param.y - param.delta
	ymax := param.y + param.delta

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{125 - contrast*n, contrast * n, contrast * n, 255}
		}
	}
	return color.RGBA{100, 100, 255, 255}
}
