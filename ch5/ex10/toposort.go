// go run toposort.go
package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, seen := topoSort(prereqs)
	errors, ok := vaildSort(order, seen, prereqs)
	if ok {
		fmt.Println("Yes it is a Topological sorting :)")
		for i, course := range order {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	} else {
		fmt.Println(errors)
	}
}

func topoSort(m map[string][]string) (order []string, seen map[string]bool) {

	var visitAll func(key string)
	seen = make(map[string]bool, 0)
	visitAll = func(key string) {
		if !seen[key] {
			seen[key] = true
			for _, item := range m[key] {
				visitAll(item)
			}
			order = append(order, key)
		}
		// for _, item := range m[key] {
		// 	if !seen[item] {
		// 		seen[item] = true
		// 		visitAll(item)
		// 		order = append(order, item)
		// 	}
		// }
	}
	for key := range m {
		visitAll(key)
	}
	return
}

func isInArray(key string, order []string) bool {
	for _, val := range order {
		if key == val {
			return true
		}
	}
	return false
}

func vaildSort(order []string, seen map[string]bool, m map[string][]string) (errorMsgs []string, ok bool) {
	// a directed graph is a linear ordering of
	// its vertices such that for every directed edge uv from vertex u to vertex v,
	// u comes before v in the ordering
	ok = true
	keyCount := 0
	for key := range seen {
		if !isInArray(key, order) {
			errorMsgs = append(errorMsgs, fmt.Sprintf("Key: %s not inclued", key))
			ok = false
			return
		}
		keyCount += 1
	}

	if len(order) != keyCount {
		errorMsgs = append(errorMsgs, "does not match key counts")
		ok = false
		return
	}
	var checkOrder func(val string, items []string)
	checkOrder = func(val string, items []string) {
		for _, item := range items {
			if val == item {
				errorMsgs = append(errorMsgs, fmt.Sprintf("Key order wrong: %s", val))
				ok = false
				return
			}
			checkOrder(val, m[item])
		}
	}
	for i := len(order) - 1; i > 0; i-- {
		checkOrder(order[i], m[order[i-1]])
	}
	return
}
