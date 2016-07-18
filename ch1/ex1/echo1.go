// echo, show the command line args
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
