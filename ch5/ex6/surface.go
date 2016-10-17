package main

import "math"
import "os"
import "fmt"

const (
    width, height = 600, 320
    cells         = 100
    xyrange       = 30.0
    xyscale       = width / 2 / xyrange
    zscale        = height * 0.4
    angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
    fmt.Fprintf(os.Stdout, "<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>\n", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay, aok := corner(i+1, j)
            bx, by, bok := corner(i, j)
            cx, cy, cok := corner(i, j+1)
            dx, dy, dok := corner(i+1, j+1)
            if !aok || !bok || !cok || !dok {
                continue
            }
            fmt.Fprintf(os.Stdout, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Fprintf(os.Stdout, "</svg>")
}

func corner(i, j int) (pointX, pointY float64, ok bool) {
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    z := f(x, y)
    if math.IsNaN(z) {
        ok = false
        return
    }
    pointX = width/2 + (x-y)*cos30*xyscale
    pointY = height/2 + (x+y)*sin30*xyscale - z*zscale
    ok = true
    return
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}
