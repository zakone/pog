// dup, show the numbers and texts which showed twice
// get texts from stdin or some files

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v \n", err)
			continue
		}
		countLines(f, counts)
		f.Close()

		for _, n := range counts {
			if n > 1 {
				fmt.Printf("dup2 of filename: %s\n", filename)
				break
			}
		}

	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
