package main

import (
	"flag"
	"fmt"
	"./tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

//type CelsiusのStringメッソド呼び出すから	
