// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 246.

// Countdown implements the countdown for a rocket launch.
package main

// NOTE: the ticker goroutine never terminates if the launch is aborted.
// This is a "goroutine leak".

import (
    "fmt"
    "os"
    "time"
)

//!+

func main() {
    // ...create abort channel...

    //!-

    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1)) // read a single byte
        abort <- struct{}{}
    }()

    //!+
    fmt.Println("Commencing countdown.  Press return to abort.")
    ticker := time.NewTicker(1 * time.Second)
    // count := 0
    // for countdown := 10; countdown > 0; countdown-- {
    //     fmt.Println(countdown)
    //     select {
    //     case <-ticker.C:
    //         count++
    //         fmt.Fprintf(os.Stdout, "Time Waiting %d Seconds", count)
    //         if count >= 10 {
    //             fmt.Fprintf(os.Stdout, "TimeOut!")
    //             ticker.Stop()
    //             return
    //         }
    //     case <-abort:
    //         fmt.Println("Launch aborted!")
    //         ticker.Stop()
    //         return
    //     }
    // }
    go func() {
        select {
        case <-ticker.C:
            count++
            fmt.Fprintf(os.Stdout, "Time Waiting %d Seconds", count)
            if count >= 10 {
                fmt.Fprintf(os.Stdout, "TimeOut!")
                ticker.Stop()
            }
        case <-abort:
            fmt.Println("Launch aborted!")
            ticker.Stop()
        }
    }()

    for input.Scan() {
        launch()
    }

}

//!-

func launch() {
    fmt.Println("Lift off!")
}
