//server of surface function
//http://localhost:8000/?width=500&height=280&cells=80&colorMax=white&colorMin=green

package main

import (
    "fmt"
    "io"
    "log"
    "math"
    "net/http"
    "strconv"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")
        getSVG(w, r)
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Params struct {
    width    int
    height   int
    cells    int
    xyrange  float64
    xyscale  float64
    zscale   float64
    angle    float64
    colorMax string
    colorMin string
}

var param Params

func getSVG(w io.Writer, r *http.Request) {

    param.width = 600
    param.height = 320
    param.cells = 100
    param.xyrange = 30.0
    param.angle = math.Pi / 6
    param.colorMax = "#ff0000"
    param.colorMin = "#0000ff"
    var sin30, cos30 = math.Sin(param.angle), math.Cos(param.angle)

    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        if k == "width" {
            param.width, _ = strconv.Atoi(v[0])
            log.Printf("change width to %d", param.width)
        } else if k == "cells" {
            param.cells, _ = strconv.Atoi(v[0])
            log.Printf("change cells to %d", param.cells)
        } else if k == "height" {
            param.height, _ = strconv.Atoi(v[0])
            log.Printf("change height to %d", param.height)
        } else if k == "xyrange" {
            param.xyrange, _ = strconv.ParseFloat(v[0], 64)
            log.Printf("change xyrange to %g", param.xyrange)
        } else if k == "colorMax" {
            param.colorMax = v[0]
        } else if k == "colorMin" {
            param.colorMin = v[0]
        }
    }
    param.xyscale = float64(param.width) / 2.0 / param.xyrange
    param.zscale = float64(param.height) * 0.4

    z := getZ(param)
    minZ, maxZ := getMinMax(z)
    middleZ := (maxZ - minZ) / 9.0
    fillColor := ""
    fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; stroke-width: 0.7' "+
        "width='%d' height='%d'>\n", param.width, param.height)
    for i := 0; i < param.cells; i++ {
        for j := 0; j < param.cells; j++ {
            ax, ay := corner(i+1, j, sin30, cos30)
            bx, by := corner(i, j, sin30, cos30)
            cx, cy := corner(i, j+1, sin30, cos30)
            dx, dy := corner(i+1, j+1, sin30, cos30)
            tempZ := z[i*param.cells+j]
            if math.IsNaN(tempZ) {
                continue
            }
            if tempZ <= middleZ {
                fillColor = param.colorMin
            } else {
                fillColor = param.colorMax
            }
            fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, fillColor)
        }
    }
    fmt.Fprintf(w, "</svg>")
}

func getMinMax(z []float64) (float64, float64) {
    min, max := z[0], z[0]
    for _, v := range z {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    return min, max
}

func getZ(param Params) []float64 {
    var z []float64
    for i := 0; i < param.cells; i++ {
        for j := 0; j < param.cells; j++ {
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
    x := param.xyrange * (float64(i)/float64(param.cells) - 0.5)
    y := param.xyrange * (float64(j)/float64(param.cells) - 0.5)
    z := calZ(x, y)
    return z
}

func corner(i int, j int, sin30 float64, cos30 float64) (float64, float64) {
    x := param.xyrange * (float64(i)/float64(param.cells) - 0.5)
    y := param.xyrange * (float64(j)/float64(param.cells) - 0.5)

    z := calZ(x, y)
    sx := float64(param.width)/2 + (x-y)*cos30*param.xyscale
    sy := float64(param.height)/2 + (x+y)*sin30*param.xyscale - z*param.zscale
    return sx, sy
}

func calZ(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}
