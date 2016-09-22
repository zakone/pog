//go run poster.go "Batman v Superman: Dawn of Justice" 2016
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	idxs := createIndexs("indexs.txt")
	baseUrl := "https://xkcd.com/%d/info.0.json"
	keywords := os.Args[1:]
	for key := range keywords {
		key, _ = strconv.Atoi(key)
		comic := idxs[key]
		url := fmt.Fprintf(baseUrl, key)
		fmt.Printf("url: %s transcript: %s\n", url, comic.Transcript)
	}
}

type Comic struct {
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Num        int    `json:"num"`
}

func createIndexs(filename string) map[int]*Comic {
	idxs := make(map[int]*Comic)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		var res Comic
		if err := json.NewDecoder([]byte(input.Text())).Decode(&res); err != nil {
			fmt.Println(err)
			continue
		}
		idxs[res.Num] = res
	}
}
