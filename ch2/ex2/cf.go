package main

import (
	"fmt"
	"os"
	"strconv"
	"./tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Foot(t)
		m := tempconv.Metre(t)
		y := tempconv.Yard(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
			f, tempconv.FToM(f), f, tempconv.FToY(f), m, tempconv.MToF(m), m, tempconv.MToY(m), y, tempconv.YToF(y), y, tempconv.YToM(y))
	}
}