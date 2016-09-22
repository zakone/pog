package main

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "os"
)

const (
    xmin, ymin, xmax, ymax = -2, -2, +2, +2
    width, height          = 1024, 1024
    eps                    = 1e-9
)

func main() {

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
    iterations := 0
    palette := []color.Color{color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}
    // answers := make([]complex128, 0, 4)
    x := real(z)
    y := imag(z)
    for {
        iterations++
        a := x*x*x - 3.0*x*y*y - 1
        b := 3.0*x*x*y - y*y*y
        c := 3.0 * (x*x - y*y)
        d := 6.0 * x * y
        x = x - (a*c+b*d)/(c*c+d*d)
        y = y - (b*c-a*d)/(c*c+d*d)
        if math.Abs((x-1)) <= eps && math.Abs(y) <= eps {
            return palette[0]
        } else if math.Abs((x + 0.5)) <= eps {
            if math.Abs(y-(math.Sqrt(3)*0.5)) <= eps {
                return palette[1]
            } else if math.Abs(y+(math.Sqrt(3)*0.5)) <= eps {
                return palette[2]
            }
        }
        if iterations >= 1000 {
            break
        }
    }
    return color.RGBA{255, 255, 255, 255}
}
