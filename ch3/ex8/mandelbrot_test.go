package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/big"
	"math"
	"os"
	"fmt"
	"testing"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func BenchmarkComplex64(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mbComplex64(z))
		}
	}
	outfile, err := os.Create("out_64.png")
    if err != nil {
        fmt.Println(err)
    }
    defer outfile.Close()
	png.Encode(outfile, img)
}

func mbComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.RGBA{125 - contrast*n, contrast * n, contrast * n, 255}
		}
	}
	return color.RGBA{100, 100, 255, 255}
}


func BenchmarkComplex128(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mbComplex128(z))
		}
	}
	outfile, err := os.Create("out_128.png")
    if err != nil {
        fmt.Println(err)
    }
    defer outfile.Close()
	png.Encode(outfile, img)
}

func mbComplex128(z complex128) color.Color {
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


func BenchmarkFloat(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		var y big.Float
		y.SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			var x big.Float
			x.SetFloat64(float64(px)/width*(xmax-xmin) + xmin)
			img.Set(px, py, mbFloat(x,y))
		}
	}
	outfile, err := os.Create("out_float.png")
    if err != nil {
        fmt.Println(err)
    }
    defer outfile.Close()
	png.Encode(outfile, img)
}

func mbFloat(x,y big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	var vx, vy big.Float
    vx.SetFloat64(0)
    vy.SetFloat64(0)
	for n := uint8(0); n < iterations; n++ {
		
		vx.Add(&tmpx, &x)
		vx = vx*vx - vy*vy + x
		vy = vx*vy + vx*vy + y
		if math.Sqrt(vx*vx + vy*vy) > 2 {
			return color.RGBA{125 - contrast*n, contrast * n, contrast * n, 255}
		}
	}
	return color.RGBA{100, 100, 255, 255}
}


func BenchmarkRat(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := big.NewRat(float64(py)/height*(ymax-ymin) + ymin, 1)
		for px := 0; px < width; px++ {
			x := big.NewRat(float64(px)/width*(xmax-xmin) + xmin, 1)
			img.Set(px, py, mbRat(*x,*y))
		}
	}
	outfile, err := os.Create("out_rat.png")
    if err != nil {
        fmt.Println(err)
    }
    defer outfile.Close()
	png.Encode(outfile, img)
}

func mbRat(x,y big.Rat) color.Color {
	const iterations = 10
	const contrast = 15

	vx := big.NewRat()
	vy := big.NewRat()
	for n := uint8(0); n < iterations; n++ {
		vx = vx*vx - vy*vy + x
		vy = vx*vy + vx*vy + y
		if math.Sqrt(vx*vx + vy*vy) > 2 {
			return color.RGBA{125 - contrast*n, contrast * n, contrast * n, 255}
		}
	}
	return color.RGBA{100, 100, 255, 255}
}


func calFloat(vx, vy, x, y big.Float) {
	var tmp1, tmp2,tmpx big.NewFloat()
	tmp1.Mul(&vx,&vx)
	tmp2.Mul(&vy,&vy)
	tmp2 = 
	tmpx.Add(*tmp1, -(*tmp2))
}

