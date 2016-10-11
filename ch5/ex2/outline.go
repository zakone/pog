// $ go build ../../ch1/ex7/fetch1.go
// $ go build outline.go
// $ ./fetch1 https://golang.org | ./outline

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "os"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "outline: %v\n", err)
        os.Exit(1)
    }
    m := make(map[string]int)
    outline(m, doc)
    for k, v := range m {
        fmt.Printf("%s\t%d\n", k, v)
    }

}

func outline(m map[string]int, n *html.Node) {
    if n.Type == html.ElementNode {
        m[n.Data] += 1
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        outline(m, c)
    }
}

// output:
// a   22
// span    3
// input   1
// textarea    2
// option  8
// html    1
// div 33
// pre 1
// br  3
// head    1
// body    1
// meta    3
// select  1
// script  10
// form    1
// iframe  1
// title   1
// link    3
