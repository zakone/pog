package main

import (
    "log"
    "net/http"
)

func main() {
    done := make(chan struct{})

    req, _ := http.NewRequest("GET", "http://www.google.com/", nil)
    req.Cancel = done

    go func() {
        //close(done)
    }()

    _, err := http.DefaultClient.Do(req)

    log.Println(err)
}