package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tempconv"
)

func main() {

	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			for _, text := range strings.Split(input.Text(), " ") {
				cf(text)
			}
		}
	} else {
		args := os.Args[1:]
		for _, arg := range args {
			cf(arg)
		}
	}

}

func cf(arg string) {
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
