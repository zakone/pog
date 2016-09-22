//go run wordFreq.go < test.txt
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {
    counts := make(map[string]int)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        word := scanner.Text()
        if match, _ := regexp.MatchString("([a-z|A-Z]+)", word); !match {
            continue
        }
        counts[word]++
    }
    fmt.Printf("word\tcount\n")
    for c, n := range counts {
        fmt.Printf("%s\t%d\n", c, n)
    }
}
