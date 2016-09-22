//go run charCount.go < test.txt
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[string]int)
	countsOhter := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsNumber(r) {
			counts["number"]++
		} else if unicode.IsSymbol(r) {
			counts["symbol"]++
		} else if unicode.IsMark(r) {
			counts["mark"]++
		} else if unicode.IsDigit(r) {
			counts["digit"]++
		} else if unicode.IsControl(r) {
			counts["control"]++
		} else if unicode.IsPunct(r) {
			counts["punct"]++
		} else {
			countsOhter[r]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	for c, n := range countsOhter {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
