// $ go run countWordsAndImages.go http://www.cnn.co.jp
package main

import (
    "bufio"
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        words, images, err := CountWordsAndImages(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "countError: %v\n", err)
            continue
        }
        fmt.Printf("words\timages\n%d\t%d\n", words, images)
    }
}

func CountWordsAndImages(url string) (words, images int, err error) {
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
    words, images = countWordsAndImages(doc)
    return
}

func countWordsAndImages(n *html.Node) (words, images int) {

    if n.Type == html.ElementNode && n.Data == "img" {
        images += 1
    }
    if n.Type == html.TextNode {
        r := strings.NewReader(n.Data)
        scanner := bufio.NewScanner(r)
        scanner.Split(bufio.ScanWords)
        for scanner.Scan() {
            words += 1
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        word, image := countWordsAndImages(c)
        words += word
        images += image
    }
    return words, images
}

// output:
// words   images
// 952 36
