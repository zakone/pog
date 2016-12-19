package title

import (
    "golang.org/x/net/html"
)

func isTitleElement(n *html.Node) bool {
    return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
    if isTitleElement(n) {
        return n.FirstChild.Data, true
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        result, ok := traverse(c)
        if ok {
            return result, true
        }
    }

    return "", false
}

func GetHtmlTitle(n *html.Node) string {
    title, ok := traverse(n)
    if !ok {
        return "notitle"
    }
    return title
}
