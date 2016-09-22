package main

import "fmt"
import "unicode"
import "bytes"

func convertSpace(b []byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, len(b)))
	if unicode.IsSpace(rune(b[0])) {
		buf.WriteByte(' ')
	} else {
		buf.WriteByte(b[0])
	}
	for i := 1; i < len(b); i++ {
		if unicode.IsSpace(rune(b[i])) {
			if unicode.IsSpace(rune(b[i-1])) {
				continue
			} else {
				buf.WriteByte(' ')
			}
		} else {
			buf.WriteByte(b[i])
		}
	}
	return buf.Bytes()
}

func main() {
	s := "g\t\n\v aa abb \f\rcccdde \tf"
	b := []byte(s)
	fmt.Printf("with space: g\\t\\n\\v aa abb \\f\\rcccdde \\tf\n")
	fmt.Printf("without space: %s\n", convertSpace(b))
}
