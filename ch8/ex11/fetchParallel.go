// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 148.

// Fetch saves the contents of a URL into a local file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)



func fetch(urls []string) (url, filename string, n int64, err error) {
	
	done := make(chan struct{})

	parallelRequest := func (url string) *http.Response {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil
		}
		req.Cancel = done
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil
		}
		return resp
	}
	
	// responses := make(chan *http.Response, len(urls))
	responses := make(chan *http.Response, len(urls))
	for _, u := range urls {
		go func(url string) {
			responses <- parallelRequest(u)
		}(u)
	}
	// respUrl := <- responses
	resp := <- responses

	// close(done)
	// loop:
	// 	for {
	// 		select {
	// 		case res, ok := <- responses:
	// 			if !ok {
	// 				continue
	// 			}
	// 			respUrl = res
	// 			close(done)
	// 			break loop
	// 		}
	// 	}

	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)

	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	close(done)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return resp.Request.URL.Host, local, n, err
	// return respUrl
}

//!-

func main() {
	url, local, n ,err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
		return
	} 
	// fmt.Printf("first response: %s \n", url)
	fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
}
