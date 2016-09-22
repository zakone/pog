package main

import (
    "github.com/nfnt/resize"
    "fmt"
    "image"
    "image/png"
    "os"
)

func main() {
    file, err := os.Open("sample.png")
    defer file.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    img = superSampling(img)
    png.Encode(os.Stdout, img)
}

func superSampling(img image.Image) image.Image {
    subPoints := []complex128{
        complex(0, 0),
        complex(0, 1),
        complex(1, 0),
        complex(1, 1),
    }
    b := img.Bounds()
    if b.Empty() {
        return nil
    }
    scaledImg := image.NewRGBA(image.Rect(0, 0, b.Dx()*2, b.Dy()*2))
    for i := 0; i < b.Dx(); i++ {
        for j := 0; j < b.Dy(); j++ {
            pointColor := img.At(i, j)
            k := complex(float64(i*2), float64(j*2))
            for _, subP := range subPoints {
                scaledImg.Set(int(real(k+subP)), int(imag(k+subP)), pointColor)
            }
        }
    }
    sampleImg := resize.Resize(uint(b.Dx()), uint(b.Dy()), scaledImg, resize.Lanczos3)
    //sampleImg := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
    //draw.Draw(sampleImg, sampleImg.Bounds(), scaledImg, image.Point{0,0}, draw.Src)
    return sampleImg
}
