//go run outline2.go http://gopl.io

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            return
        }
        doc, err := html.Parse(resp.Body)
        resp.Body.Close()
        if err != nil {
            err = fmt.Errorf("", err)
            return
        }

        var depth int
        var startElement = func(n *html.Node) {
            if n.Type == html.ElementNode {
                fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
                depth++
            }
        }
        var endElement = func(n *html.Node) {
            if n.Type == html.ElementNode {
                depth--
                fmt.Fprintf(os.Stdout, "%*s</%s>\n", depth*2, "", n.Data)
            }
        }
        forEachNode(doc, startElement, endElement)
    }
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

    if pre != nil {
        pre(n)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }

    if post != nil {
        post(n)
    }
}
