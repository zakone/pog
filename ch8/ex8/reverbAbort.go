package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    // "os"
    "strings"
    "time"
)

func echo(c *net.TCPConn, shout string, reset chan<- struct{}) {
    reset <- struct{}{}
    fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    time.Sleep(1 * time.Second)
    fmt.Fprintln(c, "\t", shout)
    time.Sleep(1 * time.Second)
    fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c *net.TCPConn) {
    input := bufio.NewScanner(c)
    resetTime := make(chan struct{})
    num := 10
    ticker := time.NewTicker(1 * time.Second)
    go func(t *time.Ticker, c *net.TCPConn) {
        fmt.Println("Start Counting Down...")
        for countdown := num; countdown > 0; countdown-- {
            select {
            case <-t.C:
                fmt.Printf("Time Left %d Seconds\n", countdown)
                if countdown == 1 {
                    fmt.Printf("TimeOut!\n")
                    c.Close()
                    t.Stop()
                }
            case <-resetTime:
                countdown = num + 1
            }
        }
    }(ticker, c)

    for input.Scan() {
        go echo(c, input.Text(), resetTime)
    }

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
