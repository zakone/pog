//go run findelement.go https://golang.org lowframe

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
    fmt.Println(n)
}

func ElementByID(doc *html.Node, id string) *html.Node {
    var target *html.Node
    search := func(n *html.Node) bool {
        if n.Type == html.ElementNode {
            for _, a := range n.Attr {
                if a.Key == "id" && a.Val == id {
                    target = n
                    return false
                }
            }
        }
        return true
    }
    forEachNode(doc, search, nil)
    fmt.Println(target)
    return target
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
    if pre != nil {
        if !pre(n) {
            return
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }

    if post != nil {
        if !post(n) {
            return
        }
    }
}

// func startElement(n *html.Node) bool {
//     if n.Type == html.ElementNode && n.Data == id {
//         return true
//     } else {
//         return false
//     }
// }

// func endElement(n *html.Node) bool {
//     if n.Type == html.ElementNode && fmt.Sprintf("/%s", n.Data) == id {
//         target = n
//         return true
//     } else {
//         return false
//     }
// }
