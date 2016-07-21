package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	done := make(chan struct{})

	req, _ := http.NewRequest("GET", "http://www.yahoo.co.jp", nil)
	req.Cancel = done

	go func() {
		//close(done)
	}()

	client := &http.Client{Timeout: time.Duration(1) * time.Microsecond}
	_, err := client.Do(req)

	log.Println(err)
}
