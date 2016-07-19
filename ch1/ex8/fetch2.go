package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	const errorURLPrefix = "http://"
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, errorURLPrefix) {
			fmt.Fprintf(os.Stderr, "Fetch: URL has error prefix %s\n", errorURLPrefix)
			continue
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		b, errs := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if errs != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
