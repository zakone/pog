package main

import (
	"io"
	"log"
	"net"
	"os"
	"flag"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		issueReport(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	var NY = flag.String("NewYork", "localhost:8010", "New York Time")
	var Tokyo = flag.String("Tokyo", "localhost:8020", "Tokyo Time")
	var London = flag.String("London", "localhost:8030", "London Time")
	locals := []string{ *NY, *Tokyo, *London }
	for _, local := range locals {
		conn, err := net.Dial("tcp", local)
		if err != nil {
		log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}

	conn, err := net.Dial("tcp", *London)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}