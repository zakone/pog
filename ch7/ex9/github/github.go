package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"sort"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//!+sortByType
type byTitle []*Issue

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byUser []*Issue

func (x byUser) Len() int           { return len(x) }
func (x byUser) Less(i, j int) bool { return x[i].User.Login < x[j].User.Login }
func (x byUser) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byNumber []*Issue

func (x byNumber) Len() int           { return len(x) }
func (x byNumber) Less(i, j int) bool { return x[i].Number < x[j].Number }
func (x byNumber) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-sortByType

func SearchIssues(terms []string, sortType string) (*IssuesSearchResult, error) {

	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if sortType == "" || sortType == "number" {
		sort.Sort(byNumber(result.Items))
	} else if sortType == "user" {
		sort.Sort(byUser(result.Items))
	} else if sortType == "title" {
		sort.Sort(byTitle(result.Items))
	}

	return &result, nil
}
