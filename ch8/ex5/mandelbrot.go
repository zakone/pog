//mandelbrot with color RGB
//go run mbColorRGBA.go > mbRGBA.png
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

func main() {
	t1 := time.Now()
	mandelbrotImg(0)
	d1 := time.Since(t1)
	fmt.Printf("In order run time: %v, with number:1 \n", d1)

	for num := 2; num <= 15; num++ {
		var wg sync.WaitGroup
		t2 := time.Now()
		for i := 0; i < num; i++ {
			wg.Add(1)
			go mandelbrotImgParallel(i, &wg)
		}
		wg.Wait()
		d2 := time.Since(t2)
		fmt.Printf("Paraller run time: %v, with goroutine number:%d \n", d2/time.Duration(num), num)
	}
	var wg sync.WaitGroup

}

// func mandelbrotImgParallel(idx int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	const (
// 		xmin, ymin, xmax, ymax = -2, -2, +2, +2
// 		width, height          = 1024, 1024
// 	)
// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	for py := 0; py < height; py++ {
// 		y := float64(py)/height*(ymax-ymin) + ymin
// 		for px := 0; px < width; px++ {
// 			x := float64(px)/width*(xmax-xmin) + xmin
// 			z := complex(x, y)
// 			img.Set(px, py, mandelbrot(z))
// 		}
// 	}
// 	output, err := os.Create(fmt.Sprintf("output%d.png", idx))
// 	if err != nil {
// 		// Openエラー処理
// 	}
// 	defer output.Close()
// 	png.Encode(output, img)
// }

func mandelbrotImg(parallel int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z, parallel))
		}
	}
	output, err := os.Create("output.png")
	if err != nil {
		// Openエラー処理
	}
	defer output.Close()
	png.Encode(output, img)
}

func mandelbrot(z complex128, parallel int) color.Color {
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
