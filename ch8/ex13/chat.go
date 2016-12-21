package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel
type clientName struct {
    c    client
    name string
}

var (
    entering = make(chan clientName)
    leaving  = make(chan client)
    messages = make(chan string) // all incoming client messages
)

func broadcaster() {
    clients := make(map[client]string) // all connected clients
    for {
        select {
        case msg := <-messages:
            // Broadcast incoming message to all
            // clients' outgoing message channels.
            for cli := range clients {
                cli <- msg
            }

        case cn := <-entering:
            for _, name := range clients {
                cn.c <- fmt.Sprintf("Ohter member: %s", name)
            }
            clients[cn.c] = cn.name

        case cli := <-leaving:
            delete(clients, cli)
            close(cli)
        }
    }
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
    ch := make(chan string) // outgoing client messages
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    ch <- "You are " + who
    messages <- who + " has arrived"

    entering <- clientName{ch, who}
    resetTime := make(chan struct{})
    minutes := 5
    ticker := time.NewTicker(60 * time.Second)
    go func() {
        fmt.Println("Start Counting Down...")
        for countdown := minutes; countdown > 0; countdown-- {
            select {
            case <-ticker.C:
                fmt.Printf("Time Left %d Minutes\n", countdown)
                if countdown == 1 {
                    fmt.Printf("TimeOut!\n")
                    conn.Close()
                    ticker.Stop()
                }
            case <-resetTime:
                countdown = minutes + 1
            }
        }
    }()
    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ": " + input.Text()
        resetTime <- struct{}{}
    }
    // NOTE: ignoring potential errors from input.Err()

    leaving <- ch
    messages <- who + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
    }
}

//!-handleConn

//!+main
func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }

    go broadcaster()
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConn(conn)
    }
}

//!-main
