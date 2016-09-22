package main

import "fmt"
import "strings"

func anagram(s1, s2 string) bool {
	s1 = removeSpace(s1)
	s2 = removeSpace(s2)
	if len(s1) != len(s2) {
		fmt.Println("diff length")
		return false
	}
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	for i := range s1 {
		if strings.Contains(s2, s1[i:i+1]) {
			s2 = removeSubS(s2, s1[i:i+1])
		} else {
			fmt.Println("not match")
			return false
		}
	}
	return true
}

func removeSpace(s string) string {
	for {
		if strings.Contains(s, " ") {
			s = removeSubS(s, " ")
		} else {
			return s
		}
	}
}
func removeSubS(s, c string) string {
	if idx := strings.Index(s, c); idx >= 0 {
		return s[:idx] + s[idx+1:]
	}
	return s
}

func main() {
	s1 := "anagrams"
	s2 := "ARS MAGNA"
	fmt.Println(s1)
	fmt.Println(s2)
	if anagram(s1, s2) {
		fmt.Println("anagram true")
	} else {
		fmt.Println("anagram false")
	}
}
