package main

import "crypto/sha256"
import "crypto/sha512"
import "fmt"
import "os"

func showSHA256(s []byte) {
    c := sha256.Sum256(s)
    fmt.Printf("sha256 value: %x", c)
}

func showSHA384(s []byte) {
    c := sha512.Sum384(s)
    fmt.Printf("sha384 value: %x", c)
}

func showSHA512(s []byte) {
    c := sha512.Sum512(s)
    fmt.Printf("sha512 value: %x", c)
}

func main() {
    params := os.Args[1:]
    if len(params) == 1 {
        showSHA256([]byte(params[0]))
    } else if params[1] == "sha384" {
        showSHA384([]byte(params[0]))
    } else if params[1] == "sha512" {
        showSHA512([]byte(params[0]))
    } else if params[1] == "sha256" {
        showSHA256([]byte(params[0]))
    }
}
