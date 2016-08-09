package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var colors = []string{
	"#0000ff",
	"#ff0000",
}

func main() {
	z := getZ()
	minZ, maxZ := getMinMax(z)
	middleZ := (maxZ - minZ) / 9.0
	svgFile := "color.svg"
	f, err := os.Create(svgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	getSVG(f, z, middleZ)
	f.Close()
}

func getSVG(file *os.File, z []float64, middleZ float64) {
	fillColor := ""
	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			tempZ := z[i*cells+j]
			if math.IsNaN(tempZ) {
				continue
			}
			if tempZ <= middleZ {
				fillColor = colors[0]
			} else {
				fillColor = colors[1]
			}
			fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, fillColor)
		}
	}
	fmt.Fprintf(file, "</svg>")
}

func getMinMax(a []float64) (float64, float64) {
	min, max := a[0], a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}

func getZ() []float64 {
	var z []float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z1 := cornerZ(i+1, j)
			z2 := cornerZ(i, j)
			z3 := cornerZ(i, j+1)
			z4 := cornerZ(i+1, j+1)
			z = append(z, (z1+z2+z3+z4)/4)
		}
	}
	return z
}

func cornerZ(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := calZ(x, y)
	return z
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := calZ(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func calZ(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

// for k := 0; k < 20; k++ {
// 	middleZ := (maxZ - minZ) / float64(k)
// 	svgFile := fmt.Sprintf("color%d.svg", k)
// 	f, err := os.Create(svgFile)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "%v\n", err)
// 	}
// 	getSVG(f, z, middleZ)
// 	f.Close()
// }
