package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	for i := 0; i < 2; i++ {
		fetchAll(i)
	}
}

func fetchAll(num int) {
	start := time.Now()
	ch := make(chan string)
	filename := fmt.Sprintf("timefetch_%d.txt", num)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("create %s file failed.", filename)
	}
	defer f.Close()
	for _, url := range os.Args[1:] {
		go fetch(url, ch, f)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, f *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while writing %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

/***
go run fetchall.go http://www.yahoo.co.jp http://edition.cnn.com https://news.google.co.jp
***/
