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

	go func() {
        close(done)
    }()
}

func fetchAll(num int) {
	start := time.Now()
	ch := make(chan string)
	done := make(chan struct{})
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	filename := fmt.Sprintf("%dtimefetch.txt", num)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("create %s file failed.", filename)
	}
	defer f.Close()
	for range os.Args[1:] {
		f.WriteString(<-ch)
		f.WriteString("\n")
	}
	str := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	f.WriteString(str)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	req, _ := http.NewRequest("GET", url, nil)
	req.Cancel = done
	//resp, err := http.Get(url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
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
