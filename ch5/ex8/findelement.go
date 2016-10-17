//go run outline2.go http://gopl.io > out.html

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "os"
)

func main() {
    resp, err := http.Get(os.Args[1])
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("", err)
        return
    }
    n := ElementByID(doc, os.Args[2])
    fmt.Println(os.Args[2])
    fmt.Println(n)
}

func ElementByID(doc *html.Node, id string) *html.Node {
    forEachNode(doc, id, startElement, endElement)
    fmt.Println(doc.Data)
    return doc
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) {
    if pre != nil {
        if pre(n, id) {
            return
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, id, pre, post)
    }

    if post != nil {
        if post(n, id) {
            return
        }
    }
}

func startElement(n *html.Node, id string) bool {
    if n.Type == html.ElementNode && n.Data == id {
        return true
    } else {
        return false
    }
}

func endElement(n *html.Node, id string) bool {
    if n.Type == html.ElementNode && fmt.Sprintf("/%s", n.Data) == id {
        return true
    } else {
        return false
    }
}
