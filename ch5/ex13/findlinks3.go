// ./findlinks3 https://golang.org
package main

import (
	"./extract"
	"./title"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(link string) []string {
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
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
		os.Mkdir(u.Host, 0777)
		if ioutil.WriteFile(fmt.Sprintf("%s/%s.txt", u.Host, title), body, 0644) != nil {
			fmt.Errorf("write file error: %s", title)
		}
	}

	for _, subLink := range list {
		subURL, err := url.Parse(subLink)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if subURL.Host != u.Host {
			continue
		}
		savePage(subLink)
	}
	return list
}

// https://golang.org
// https://golang.org/
// https://golang.org/doc/
// https://golang.org/pkg/
// https://golang.org/project/
// https://golang.org/help/
// https://golang.org/blog/
// http://play.golang.org/
// https://tour.golang.org/
// https://golang.org/dl/
// https://blog.golang.org/
// https://developers.google.com/site-policies#restrictions
// https://golang.org/LICENSE
// https://golang.org/doc/tos.html
// http://www.google.com/intl/en/policies/privacy/
