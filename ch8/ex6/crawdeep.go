// ./crawdeep -depth=3 https://golang.org
package main

import (
	"fmt"
	"log"
	"flag"
	"sync"
	"./links"
)


var tokens = make(chan struct{}, 20)

func crawl(url string, wg *sync.WaitGroup) []string {
	defer wg.Done()
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

var depth = flag.Int("depth", 3, "crawl depth")

//!+
func main() {
	flag.Parse()
	worklist := make(chan []string)
	urls := flag.Args()
	go func() { worklist <- urls }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	var wg sync.WaitGroup
	for d := 0; d < *depth; d++ {
		list := <-worklist
		fmt.Printf("Crawl in Depth %d\n", d)
		fmt.Println(list)
		for _, link := range list {
			wg.Add(1)
			if !seen[link] {
				seen[link] = true
				go func(link string, wg *sync.WaitGroup) {
					worklist <- crawl(link, wg)
				}(link, &wg)
			}
		}
	}
	wg.Wait()

}

//!-
