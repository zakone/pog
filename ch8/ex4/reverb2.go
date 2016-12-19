// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"sync"
)

func echo(c *net.TCPConn, shout string, wg *sync.WaitGroup, nums chan<- int64) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(1*time.Second)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(1*time.Second)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	nums <- 1
	// nums chan<- int64
}

//!+
func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	nums := make(chan int64)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), &wg, nums)
	}
	go func(){
		wg.Wait()
		c.CloseWrite()
		close(nums)
	}()
	var total int64
	for num := range nums {
		total += num
	}
	fmt.Printf("Echo number: %d\n", total)
}

//!-

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
    }
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
		
	}
}
