package main

import "fmt"

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("KB: %f \n", KB)
	fmt.Printf("MB: %f \n", MB)
	fmt.Printf("GB: %f \n", GB)
	fmt.Printf("TB: %f \n", TB)
	fmt.Printf("PB: %f \n", PB)
	fmt.Printf("EB: %f \n", EB)
	fmt.Printf("ZB: %f \n", ZB)
	fmt.Printf("YB: %f \n", YB)
}
