
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			for _, elem := range tok.Attr {
				stack = append(stack, fmt.Sprintf("%s=%s", elem.Name.Local, elem.Value))
			}
		case xml.EndElement:
			idx := stackIndex(stack, tok.Name.Local)
			if idx == -1 {
				fmt.Fprintf(os.Stderr, "find index error")
				os.Exit(1)
			}
			stack = stack[:idx]  // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func stackIndex(stack []string, s string) int{
	for i := len(stack)-1; i >= 0; i-- {
		if stack[i] == s {
			return i
		}
	}
	return -1
}

// ./fetch1 http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlparse div class=head h2
//!-