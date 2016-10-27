// go run fetch.go https://golang.org/doc/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	filename, _, err := fetch(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(filename)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()
	n, errs := io.Copy(f, resp.Body)
	if errs != nil {
		return "", 0, errs
	}
	return local, n, errs
}
