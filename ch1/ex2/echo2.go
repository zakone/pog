// echo, show the command line args
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("index: %d, arg: %s\n", index, arg)
	}
}
