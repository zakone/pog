package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {

	var palette = []color.Color{color.RGBA{0, 0, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{128, 128, 0, 255},
		color.RGBA{0, 128, 128, 255},
		color.RGBA{128, 0, 128, 255}}

	type Params struct {
		cycles  float64
		res     float64
		size    int
		nframes int
		delay   int
	}
	param := Params{5.0, 0.001, 100, 64, 8}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			param.cycles, _ = strconv.ParseFloat(v[0], 64)
			log.Printf("change cycles to %f", param.cycles)
		} else if k == "res" {
			param.res, _ = strconv.ParseFloat(v[0], 64)
			log.Printf("change res to %f", param.res)
		} else if k == "size" {
			param.size, _ = strconv.Atoi(v[0])
			log.Printf("change size to %d", param.size)
		} else if k == "delay" {
			param.delay, _ = strconv.Atoi(v[0])
			log.Printf("change delay to %d", param.delay)
		} else if k == "nframes" {
			param.nframes, _ = strconv.Atoi(v[0])
			log.Printf("change nframes to %d", param.nframes)
		}
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: param.nframes}
	phase := 0.0
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < param.nframes; i++ {
		rect := image.Rect(0, 0, 2*param.size+1, 2*param.size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := uint8(rand.Intn(len(palette)-1) + 1)
		for t := 0.0; t < param.cycles*2*math.Pi; t += param.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(param.size+int(x*float64(param.size)+0.5), param.size+int(y*float64(param.size)+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, param.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
