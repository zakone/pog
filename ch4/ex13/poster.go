package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const PosterURL = "http://www.omdbapi.com/?"

func main() {
	title := os.Args[1]
	year := os.Args[2]
	poster(title, year)

}

type Movie struct {
	Poster string `json:"Poster"`
}

func poster(title string, year string) {
	resp, err := http.Get(PosterURL + "t=" + url.QueryEscape(title) + "&y=" + year)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// bufbody := new(bytes.Buffer)
	// bufbody.ReadFrom(resp.Body)
	if resp.StatusCode != http.StatusOK {
		// fmt.Println(bufbody.String())
		fmt.Println("search query failed: ", resp.Status)
		os.Exit(1)
	}

	var val Movie
	if err := json.NewDecoder(resp.Body).Decode(&val); err != nil {
		fmt.Println("JSON Unmarshal failed: %s", err)
		os.Exit(1)
	}
	response, err := http.Get(val.Poster)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("search query failed: %s", response.Status)
		os.Exit(1)
	}
	img, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filename := title + year + ".jpg"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	f.Write(img)
}
