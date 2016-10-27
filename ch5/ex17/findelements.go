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
    nodes := ElementsByTagName(doc, os.Args[2:]...)
    for _, node := range nodes {
        fmt.Println(node.Data)
    }
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
    var nodes []*html.Node
    for _, name := range names {
        fmt.Println(name)
        res := forEachNode(nil, doc, name, startElement, nil)
        nodes = append(nodes, res...)
    }
    return nodes
}

func forEachNode(nodes []*html.Node, n *html.Node, name string, pre, post func(nodes []*html.Node, n *html.Node, name string)) []*html.Node {
    if pre != nil {
        pre(nodes, n, name)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(nodes, c, name, pre, post)
    }

    if post != nil {
        post(nodes, n, name)
    }
    return nodes
}

func startElement(nodes []*html.Node, n *html.Node, name string) {
    if n.Type == html.ElementNode && n.Data == name {
        nodes = append(nodes, n)
    }
}
