package main

import (
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

    superImg := image.NewRGBA(image.Rect(0, 0, img.Width*2, img.Height*2))
    for i := 0; i < img.Width; i++ {
        for j := 0; j < img.Height; j++ {
            pointColor := img.At(i, j)
            k := complex(float64(i*2), float64(j*2))
            for subP := range subPoints {
                superImg.Set(real(k+subP), imag(k+subP), pointColor)
            }
        }
    }
    return superImg
}
