// $ go run htmlparse.go https://golang.org

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "io"
    "io/ioutil"
    "net/http"
    "os"
)

type HtmlReader struct {
    body []byte
    i    int64
    io.Reader
}

func (h *HtmlReader) Read(b []byte) (n int, err error) {
    if h.i >= int64(len(h.body)) {
        return 0, io.EOF
    }
    n = copy(b, h.body[h.i:])
    h.i += int64(n)
    return
}

func NewReader(s string) *HtmlReader {
    var h HtmlReader
    resp, err := http.Get(s)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
        os.Exit(1)
    }
    h.body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Copy: %v\n", err)
        os.Exit(1)
    }
    resp.Body.Close()
    return &h
}

func main() {

    h := NewReader(os.Args[1])
    doc, err := html.Parse(h)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Parse: %v\n", err)
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
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }
    return links
}
