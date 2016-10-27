package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {

	defer func() {
		switch p := recover(); p.(type) {
		case nil:
			// sorry nothing
		case string:
			pi, ok := p.(string)
			if !ok {
				break
			}
			if !strings.Contains(pi, "Return:") {
				break
			}
			host := strings.Split(pi, ":")[1]
			fmt.Printf("Host name: %s\n", host)
			break
		default:
			panic(p)
		}
	}()

	nonzero(os.Args[1])
}

func nonzero(link string) {
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	if u.Host != "" {
		panic("Return:" + u.Host)
	}
}
