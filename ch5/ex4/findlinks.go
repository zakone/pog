// $ go build ../../ch1/ex7/fetch1.go
// $ go build findlinks.go
// $ ./fetch1 https://golang.org | ./findlinks

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "os"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
        os.Exit(1)
    }
    visit(nil, doc)
}

func visit(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                links = append(links, a.Val)
                fmt.Println(a.Val)
            }
        }
    } else if n.Type == html.ElementNode && n.Data == "img" {
        for _, a := range n.Attr {
            if a.Key == "src" {
                links = append(links, a.Val)
                fmt.Println(a.Val)
            }
        }
    } else if n.Type == html.ElementNode && n.Data == "script" {
        for _, a := range n.Attr {
            if a.Key == "src" {
                links = append(links, a.Val)
                fmt.Println(a.Val)
            }
        }
    } else if n.Type == html.ElementNode && n.Data == "link" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                links = append(links, a.Val)
                fmt.Println(a.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }
    return links
}
