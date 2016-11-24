package xmltree

import "encoding/xml"

type Node interface{}

type CharData string

type Element struct {
	Type xml.Name
	Attr []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	setRoot(dec xml.NewDecoder)
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

func setRoot(dec xml.NewDecoder) {
	tok, err := dec.Token()
	if err == io.EOF {
		return
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
	case xml.CharData:
		if containsAll(stack, os.Args[1:]) {
			fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
		}
	}
}