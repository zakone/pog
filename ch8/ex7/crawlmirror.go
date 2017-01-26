package main

import (
	"./links"
	"./title"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

var tokens = make(chan struct{}, 20)

func crawl(link, host string, wg *sync.WaitGroup) []string {

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(link)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	var savePage = func(link string) {
		resp, err := http.Get(link)
		defer resp.Body.Close()
		if err != nil {
			fmt.Errorf("Get link %s: %s", link, err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Errorf("getting %s: %s", link, resp.Status)
			return
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Errorf("parsing %s as HTML: %v", link, err)
			return
		}
		title := title.GetHtmlTitle(doc)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Errorf("reading resp body to bytes: %v", err)
			return
		}
		os.Mkdir(host, 0777)
		if ioutil.WriteFile(fmt.Sprintf("%s/%s.txt", host, title), body, 0644) != nil {
			fmt.Errorf("write file error: %s", title)
		}
	}
	for _, subLink := range list {
		subURL, err := url.Parse(subLink)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if subURL.Host != host {
			continue
		}
		savePage(subLink)
	}
	return list
}

var depth = flag.Int("depth", 3, "crawl depth")

//!+
func main() {
	flag.Parse()
	urls := flag.Args()
	worklist := make(chan []string)
	url, _ := url.Parse(urls[0])

	go func() { worklist <- urls }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	var wg sync.WaitGroup
	for d := 0; d < *depth; d++ {
		list := <-worklist
		fmt.Printf("Crawl in Depth %d\n", d)
		for _, link := range list {
			wg.Add(1)
			if !seen[link] {
				seen[link] = true
				go func(link, u string, wg *sync.WaitGroup) {
					worklist <- crawl(link, u, wg)
				}(link, url.Host, &wg)
			}
		}
	}
	wg.Wait()

}

//!-
