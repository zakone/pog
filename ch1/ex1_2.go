// echo, show the command line args
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println("index: %d, arg: %s", index, arg)
	}
}
