//go run findelement.go https://golang.org a

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
    fmt.Println(n.Data)
}

func ElementByID(doc *html.Node, id string) *html.Node {
    var target *html.Node
    target = doc
    forEachNode(doc, id, target, startElement, endElement)
    fmt.Println(target.Data)
    return target
}

func forEachNode(n *html.Node, id string, target *html.Node, pre, post func(n *html.Node, target *html.Node, id string) bool) {
    if pre != nil {
        if pre(n, target, id) {
            return
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, id, target, pre, post)
    }

    if post != nil {
        if post(n, target, id) {
            return
        }
    }
}

func startElement(n *html.Node, target *html.Node, id string) bool {
    if n.Type == html.ElementNode && n.Data == id {
        target = n
        return true
    } else {
        return false
    }
}

func endElement(n *html.Node, target *html.Node, id string) bool {
    if n.Type == html.ElementNode && fmt.Sprintf("/%s", n.Data) == id {
        target = n
        return true
    } else {
        return false
    }
}
