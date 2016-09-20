package main

import "crypto/sha256"
import "fmt"

func countSHA(str1, str2 []byte) int {
	c1 := sha256.Sum256(str1)
	c2 := sha256.Sum256(str2)
	diffCount := 0
	for i := range c1 {
		if c1[i] != c2[i] {
			diffCount += 1
		}
	}
	return diffCount
}

func main() {
	fmt.Println(countSHA([]byte("x"), []byte("X")))
}
