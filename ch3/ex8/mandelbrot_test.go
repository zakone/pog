// BenchmarkComplex64-4 	       1	1050163132 ns/op
// BenchmarkComplex128-4	2000000000	         0.27 ns/op
// BenchmarkFloat-4     	       1	81788334177 ns/op
// BenchmarkRat-4       	       1	167398292570 ns/op

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
	"testing"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	eps                    = 1e-9
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
		if float32(cmplx.Abs(complex128(v))) > 2 {
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
		y := new(big.Float).SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := new(big.Float).SetFloat64(float64(px)/width*(xmax-xmin) + xmin)
			img.Set(px, py, mbFloat(x, y))
		}
	}
	outfile, err := os.Create("out_float.png")
	if err != nil {
		fmt.Println(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}

func mbFloat(x, y *big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	vx := new(big.Float).SetFloat64(0.0)
	vy := new(big.Float).SetFloat64(0.0)
	for n := uint8(0); n < iterations; n++ {
		vMulVPlusZFloat(vx, vy, x, y)
		if sqrtFloat(vAbsFloat(vx, vy)).Cmp(new(big.Float).SetInt64(2)) > 0 {
			return color.RGBA{125 - contrast*n, contrast * n, contrast * n, 255}
		}
	}
	return color.RGBA{100, 100, 255, 255}
}

func BenchmarkRat(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := new(big.Rat).SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := new(big.Rat).SetFloat64(float64(px)/width*(xmax-xmin) + xmin)
			img.Set(px, py, mbRat(x, y))
		}
	}
	outfile, err := os.Create("out_rat.png")
	if err != nil {
		fmt.Println(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}

func mbRat(x, y *big.Rat) color.Color {
	const iterations = 5
	const contrast = 15

	vx := new(big.Rat).SetFloat64(0)
	vy := new(big.Rat).SetFloat64(0)
	for n := uint8(0); n < iterations; n++ {
		vMulVPlusZRat(vx, vy, x, y)
		if sqrtRat(vAbsRat(vx, vy)).Cmp(new(big.Rat).SetInt64(2)) > 0 {
			return color.RGBA{125 - contrast*n, contrast * n, contrast * n, 255}
		}
	}
	return color.RGBA{100, 100, 255, 255}
}

func vMulVPlusZFloat(vx, vy, x, y *big.Float) {
	//vx = vx*vx - vy*vy + x
	realMul := new(big.Float).Add(new(big.Float).Mul(vx, vx), new(big.Float).Neg(new(big.Float).Mul(vy, vy)))
	//vy = vx*vy + vx*vy + y
	imagMul := new(big.Float).Add(new(big.Float).Mul(vx, vy), new(big.Float).Mul(vx, vy))
	vx.Add(realMul, x)
	vy.Add(imagMul, y)
}

func vAbsFloat(vx, vy *big.Float) *big.Float {
	tmp := new(big.Float).Add(new(big.Float).Mul(vx, vx), new(big.Float).Mul(vy, vy))
	return tmp
}

func vMulVPlusZRat(vx, vy, x, y *big.Rat) {
	//vx = vx*vx - vy*vy + x
	realMul := new(big.Rat).Add(new(big.Rat).Mul(vx, vx), new(big.Rat).Neg(new(big.Rat).Mul(vy, vy)))
	//vy = vx*vy + vx*vy + y
	imagMul := new(big.Rat).Add(new(big.Rat).Mul(vx, vy), new(big.Rat).Mul(vx, vy))
	vx.Add(realMul, x)
	vy.Add(imagMul, y)
}

func vAbsRat(vx, vy *big.Rat) *big.Rat {
	tmp := new(big.Rat).Add(new(big.Rat).Mul(vx, vx), new(big.Rat).Mul(vy, vy))
	return tmp
}

// func sqrtFloat32(val float32) float32 {
// 	x := float32(1.0)
// 	p := x
// 	for {
// 		x = 1 / 2 * (x + (val / x))
// 		minus := x - p
// 		if minus < 0 {
// 			minus = -minus
// 		}
// 		if minus < eps {
// 			break
// 		}
// 		p = x
// 	}
// 	return x
// }

func sqrtFloat(val *big.Float) *big.Float {
	half := new(big.Float).SetFloat64(0.5)
	x := new(big.Float).SetInt64(1)
	p := new(big.Float).Neg(x)
	t := new(big.Float)
	epsFloat := new(big.Float).SetFloat64(eps)
	for {
		t.Quo(val, x)
		t.Add(x, t)
		x.Mul(half, t)
		if new(big.Float).Add(x, p).Cmp(epsFloat) < 0 {
			break
		}
		p.Neg(x)
	}
	return x
}

func sqrtRat(val *big.Rat) *big.Rat {
	half := new(big.Rat).SetFloat64(0.5)
	x := new(big.Rat).SetInt64(1)
	p := new(big.Rat).Neg(x)
	t := new(big.Rat)
	epsFloat := new(big.Rat).SetFloat64(eps)
	for {
		t.Quo(val, x)
		t.Add(x, t)
		x.Mul(half, t)
		if new(big.Rat).Add(x, p).Cmp(epsFloat) < 0 {
			break
		}
		p.Neg(x)
	}
	return x
}
