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
    "linear algebra":        {"calculus"}, //unvaild key pair
}

func main() {
    order, errors, ok := topoSort(prereqs)
    if ok {
        fmt.Println("Yes it is a Topological sorting :)")
        for i, course := range order {
            fmt.Printf("%d:\t%s\n", i+1, course)
        }
    } else {
        for _, err := range errors {
            fmt.Println(err)
        }
    }
}

func topoSort(m map[string][]string) (order, errorMsgs []string, ok bool) {

    ok = true
    // Check cyclulation
    var checkCycle func(key string, items []string)
    checkCycle = func(key string, items []string) {
        for _, item := range items {
            if isInArray(key, m[item]) {
                errorMsgs = append(errorMsgs, fmt.Sprintf("cyclulation wrong. Key: %s\n", key))
                ok = false
                return
            }
            checkCycle(key, m[item])
        }
    }
    // Topological sort
    var visitAll func(key string)
    seen := make(map[string]bool, 0)
    visitAll = func(key string) {
        if !seen[key] {
            seen[key] = true
            for _, item := range m[key] {
                visitAll(item)
            }
            order = append(order, key)
        }
    }
    for key := range m {
        checkCycle(key, m[key])
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
