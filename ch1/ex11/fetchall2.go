package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchAll()
}

func fetchAll() {
	start := time.Now()
	ch := make(chan string)
	done := make(chan struct{})
	for _, url := range os.Args[1:] {
		go fetch(url, ch, done)
	}
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	for index := range os.Args[1:] {
		select {
		case <-done:
			fmt.Printf("request%d cancled!\n", index)
			continue
		case <-ch:
			fmt.Println(<-ch)
		}
	}
	fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, done <-chan struct{}) {
	start := time.Now()
	req, _ := http.NewRequest("GET", url, nil)
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	// client := &http.Client{Timeout: time.Duration(5) * time.Second}
	// resp, err := client.Do(req)

	if err != nil {
		//ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		//ch <- fmt.Sprintf("while writing %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	//ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

/***
go run fetchall2.go http://www.yahoo.co.jp http://edition.cnn.com https://news.google.co.jp
***/
