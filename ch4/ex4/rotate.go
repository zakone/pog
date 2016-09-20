package main

import "fmt"

func rotate(s []int, dir string, left, right int) []int{
	tmp := make([]int, 0, len(s))
	tmp = append(tmp, s[:left]...)
	tmp = append(tmp, s[right:]...)
	if dir == "right" {
		return append(s[left:right], tmp...)
	} else if dir == "left" {
		return append(tmp, s[left:right]...)
	} else {
		return s
	}
}

func main() {
	a := []int{0,1,2,3,4,5}
	a = rotate(a,"right",2,4)
	fmt.Println(a)
}