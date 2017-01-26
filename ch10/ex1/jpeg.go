package main

import (
    "flag"
    "fmt"
    "image"
    "image/gif"
    "image/jpeg"
    _ "image/png" // register PNG decoder
    "io"
    "os"
)

var imageFormat = flag.String("imageFormat", "jpeg", "Enter Output Image Format")

func main() {
    flag.Parse()
    if err := toJPEG(os.Stdin, os.Stdout); err != nil {
        fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
        os.Exit(1)
    }
}

func toJPEG(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
    if err != nil {
        return err
    }
    fmt.Fprintln(os.Stderr, "Input format =", kind)
    if *imageFormat == "jpeg" {
        return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
    } else if *imageFormat == "gif" {
        return gif.Encode(out, img, &jpeg.Options{Quality: 95})
    } else if *imageFormat == "png" {
        return png.Encode(out, img, &jpeg.Options{Quality: 95})
    } else {
        fmt.Fprint("Unsupported Format")
        os.Exit(1)
    }

}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
