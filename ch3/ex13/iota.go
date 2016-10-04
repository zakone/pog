package main

import "fmt"

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
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
