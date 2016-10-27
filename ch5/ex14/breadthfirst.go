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
    var keys []string
    for key := range prereqs {
        keys = append(keys, key)
    }
    breadthFirst(topoSort, keys)
}

func breadthFirst(f func(item string, checked map[string]bool) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                worklist = append(worklist, f(item, seen)...)
            }
        }
    }
}

func topoSort(key string, checked map[string]bool) []string {
    var visitAll func(key string)
    visitAll = func(key string) {
        if !checked[key] {
            checked[key] = true
            for _, item := range prereqs[key] {
                visitAll(item)
            }
            fmt.Println(key)
        }
    }
    visitAll(key)
    return prereqs[key]
}
