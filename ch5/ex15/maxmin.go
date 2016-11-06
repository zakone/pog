package maxmin

import (
	"fmt"
)

// must have one, else call error
func Max2(val int, vals ...int) int {
	largest := val
	for _, v := range vals {
		if v >= largest {
			largest = v
		}
	}
	return largest
}

// ok to call no params
func MaxOne(vals ...int) (int, bool) {
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

// should be this way!
// but be care of val & vals
func Min2(val int, vals ...int) int {
	minest := val
	for _, v := range vals {
		if v <= minest {
			minest = v
		}
	}
	return minest
}

func MinOne(vals ...int) (int, bool) {
	if len(vals) == 0 {
		fmt.Println("Error! Must input one number")
		return 0, false
	}
	minest := vals[0]
	for _, v := range vals {
		if v < minest {
			minest = v
		}
	}
	return minest, true
}
