//go run outline2.go http://gopl.io > out.html

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "os"
    "strings"
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

var depth int

func startElement(n *html.Node) {
    if n.Type == html.ElementNode {
        fmt.Printf("%*s<%s", depth*2, "", n.Data)
        for _, a := range n.Attr {
            fmt.Fprintf(os.Stdout, " %s=%q", a.Key, a.Val)
        }
        if n.FirstChild != nil {
            if n.FirstChild.Type == html.ElementNode || n.FirstChild.NextSibling != nil || (n.Data == "script" || n.Data == "style") {
                fmt.Fprintf(os.Stdout, ">\n")
                depth++
            } else {
                fmt.Fprintf(os.Stdout, ">")
            }

        } else {
            fmt.Fprintf(os.Stdout, "/>\n")
        }
    } else if n.Type == html.TextNode || n.Type == html.CommentNode {
        s := strings.TrimSpace(n.Data)

        if s != "" {
            if n.NextSibling != nil || (n.PrevSibling != nil && n.PrevSibling.Data == "a") || (n.Parent != nil && (n.Parent.Data == "script" || n.Parent.Data == "style")) {
                for _, i := range strings.Split(s, "\n") {
                    fmt.Fprintf(os.Stdout, "%*s%s\n", depth*2, "", i)
                }
            } else {
                fmt.Fprintf(os.Stdout, "%s", s)
            }
        }
    }
}

func endElement(n *html.Node) {
    if n.Type == html.ElementNode && n.FirstChild != nil {
        if n.FirstChild.NextSibling != nil || n.FirstChild.Type == html.ElementNode || (n.Data == "script" || n.Data == "style") {
            depth--
            fmt.Fprintf(os.Stdout, "%*s</%s>\n", depth*2, "", n.Data)
        } else {
            fmt.Fprintf(os.Stdout, "</%s>\n", n.Data)
        }

    }
}
