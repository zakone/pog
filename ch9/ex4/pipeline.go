package main

import "fmt"

//!+
func start(out chan<- int) {
	out <- x
	// close(out)
}

func pipeline(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v
	}
	// close(out)
}

func end(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go start(in)
	go squarer(out, in)
	end(out)
}

//!-
