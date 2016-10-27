package main

import (
	"fmt"
)

func main() {
	fmt.Println(join("a", "b", "c"))
}
func join(vals ...string) string {
	res := ""
	for _, val := range vals {
		res = res + val
	}
	return res
}
