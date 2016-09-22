package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const ComicURL = "https://xkcd.com/%d/info.0.json"
const MaxIndexs = 5000

type Comic struct {
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Num        int    `json:"num"`
}

func scrap(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	i := 0
	for i < MaxIndexs {
		url := fmt.Sprintf(ComicURL, i)
		fmt.Println(url)
		i++
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("request error: %s\n", err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Errorf("search query failed: %s", resp.Status)
			continue
		}
		var result Comic
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Printf("json decoder error: %s\n", err)
			continue
		}
		data, err := json.Marshal(result)
		if err != nil {
			fmt.Printf("JSON marshaling failed: %s\n", err)
			continue
		}
		f.Write(data)
		f.WriteString("\n")

	}
}

func main() {
	filename := os.Args[1]
	scrap(filename)
}
