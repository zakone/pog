//go run comicSearch.go 6
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	idxs := createIndexs("indexs.txt")
	keywords := os.Args[1:]
	for _, key := range keywords {
		key, _ := strconv.Atoi(key)
		comic := idxs[key]
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", key)
		fmt.Printf("url: %s transcript: %s\n", url, comic.Transcript)
	}
}

type Comic struct {
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Num        int    `json:"num"`
}

func createIndexs(filename string) map[int]Comic {
	idxs := make(map[int]Comic)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		var res Comic
		reader := strings.NewReader(input.Text())
		if err := json.NewDecoder(reader).Decode(&res); err != nil {
			fmt.Println(err)
			continue
		}
		idxs[res.Num] = res
	}

	return idxs
}
