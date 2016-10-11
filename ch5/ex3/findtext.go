// $ go build ../../ch1/ex7/fetch1.go
// $ go build findtext.go
// $ ./fetch1 https://golang.org | ./findtext

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "os"
    "strings"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "outline: %v\n", err)
        os.Exit(1)
    }
    visit(doc)
}

func visit(n *html.Node) {

    if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
        //go to sibling
        for c := n.NextSibling; c != nil; c = c.NextSibling {
            visit(c)
        }
    } else {
        //go to child
        if n.Type == html.TextNode {
            //remove newline
            s := strings.TrimSpace(n.Data)
            if s != "" {
                fmt.Printf("textNode: %q is outputed.\n", s)
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            visit(c)
        }
    }
}

//output:
// textNode: "The Go Programming Language" is outputed.
// textNode: "..." is outputed.
// textNode: "The Go Programming Language" is outputed.
// textNode: "Go" is outputed.
// textNode: "▽" is outputed.
// textNode: "Documents" is outputed.
// textNode: "Packages" is outputed.
// textNode: "The Project" is outputed.
// textNode: "Help" is outputed.
// textNode: "Blog" is outputed.
// textNode: "Play" is outputed.
// textNode: "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, 世界\")\n}" is outputed.
// textNode: "Run" is outputed.
// textNode: "Format" is outputed.
// textNode: "Share" is outputed.
// textNode: "Pop-out" is outputed.
// textNode: "Try Go" is outputed.
// textNode: "// You can edit this code!\n// Click here and start typing.\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, 世界\")\n}" is outputed.
// textNode: "Hello, 世界" is outputed.
// textNode: "Run" is outputed.
// textNode: "Share" is outputed.
// textNode: "Tour" is outputed.
// textNode: "Hello, World!" is outputed.
// textNode: "Conway's Game of Life" is outputed.
// textNode: "Fibonacci Closure" is outputed.
// textNode: "Peano Integers" is outputed.
// textNode: "Concurrent pi" is outputed.
// textNode: "Concurrent Prime Sieve" is outputed.
// textNode: "Peg Solitaire Solver" is outputed.
// textNode: "Tree Comparison" is outputed.
// textNode: "Go is an open source programming language that makes it easy to build\nsimple, reliable, and efficient software." is outputed.
// textNode: "Download Go" is outputed.
// textNode: "Binary distributions available for" is outputed.
// textNode: "Linux, Mac OS X, Windows, and more." is outputed.
// textNode: "Featured video" is outputed.
// textNode: "Featured articles" is outputed.
// textNode: "Read more" is outputed.
// textNode: "Build version go1.7.1." is outputed.
// textNode: "Except as" is outputed.
// textNode: "noted" is outputed.
// textNode: ",\nthe content of this page is licensed under the\nCreative Commons Attribution 3.0 License,\nand code is licensed under a" is outputed.
// textNode: "BSD license" is outputed.
// textNode: "." is outputed.
// textNode: "Terms of Service" is outputed.
// textNode: "|" is outputed.
// textNode: "Privacy Policy" is outputed.
// textNode: "Build version go1.7.1." is outputed.
// textNode: "Except as" is outputed.
// textNode: "noted" is outputed.
// textNode: ",\nthe content of this page is licensed under the\nCreative Commons Attribution 3.0 License,\nand code is licensed under a" is outputed.
// textNode: "BSD license" is outputed.
// textNode: "." is outputed.
// textNode: "Terms of Service" is outputed.
// textNode: "|" is outputed.
// textNode: "Privacy Policy" is outputed.
