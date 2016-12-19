package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	"fmt"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

var port = flag.Int("port", 8000, "server port")

func main() {
	flag.Parse()
	host := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}