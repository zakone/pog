package main

import (
	"fmt"
)

// func max(vals ...int) int {
// 	largest := 0
// 	for _, v := range vals {
// 		if v >= largest {
// 			largest = v
// 		}
// 	}
// 	return largest
// }

func maxOne(vals ...int) (int, bool) {
	if len(vals) == 0 {
		fmt.Println("Error! Must input one number")
		return 0, false
	}
	largest := vals[0]
	for _, v := range vals {
		if v > largest {
			largest = v
		}
	}
	return largest, true
}

// func min(vals ...int) int {
// 	minest := vals[0]
// 	for _, v := range vals {
// 		if v <= minest {
// 			largest = v
// 		}
// 	}
// 	return largest
// }

func minOne(vals ...int) (int, bool) {
	if len(vals) == 0 {
		fmt.Println("Error! Must input one number")
		return 0, false
	}
	minest := vals[0]
	for _, v := range vals {
		if v < minest {
			largest = v
		}
	}
	return minest, true
}
