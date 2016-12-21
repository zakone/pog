package main

import (
	"fmt"
	"log"
	"os"
	"./links"
)

var tokens = make(chan struct{}, 20)
var done = make(chan struct{})

func crawl(url string) []string {
	fmt.Println(url)

	select {
	case tokens <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-tokens }()
	list, err := links.Extract(url, done)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {

	worklist := make(chan []string)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
//!-
