package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var NY = flag.String("NewYork", "localhost:8010", "New York Time")
var Tokyo = flag.String("Tokyo", "localhost:8020", "Tokyo Time")
var London = flag.String("London", "localhost:8030", "London Time")

func main() {

	locals := []string{*NY, *Tokyo}

	for _, local := range locals {
		go func() {
			tcpAddr, err := net.ResolveTCPAddr("tcp", local)
			if err != nil {
				println("ResolveTCPAddr failed:", err.Error())
				os.Exit(1)
			}
			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				log.Fatal(err)
			}
			done := make(chan struct{})
			go func() {
				io.Copy(os.Stdout, conn) // NOTE: ignoring errors
				log.Println("done")
				done <- struct{}{} // signal the main goroutine
			}()
			mustCopy(conn, os.Stdin)
			conn.CloseWrite()
			<-done // wait for background goroutine to finish
		}()

	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", *London)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
