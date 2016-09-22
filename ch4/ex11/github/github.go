package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const IssueURL = "https://api.github.com/repos/zakone/pog/issues"

type Issue struct {
	Number    int       `json:"number,omitempty"`
	HTMLURL   string    `json:"html_url"`
	Title     string    `json:"title,omitempty"`
	State     string    `json:"state,omitempty"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body,omitempty"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
}

func CreateIssue(title string, body string) (*Issue, error) {

	var issueReq = IssueRequest{Title: title, Body: body}
	data, err := json.Marshal(issueReq)
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s\n", err)
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		IssueURL,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("d04863358ea7c703b6b1d8dffe1ffae3126b31dd", "x-oauth-basic")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetIssues() ([]Issue, error) {

	resp, err := http.Get(IssueURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetSingleIssue(number string) (*Issue, error) {

	resp, err := http.Get(IssueURL + "/" + number)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func EditIssue(number string, title string) (*Issue, error) {

	var editReq = IssueRequest{Title: title}
	data, err := json.Marshal(editReq)
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s\n", err)
		return nil, err
	}
	req, err := http.NewRequest(
		"PATCH",
		IssueURL+"/"+number,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("d04863358ea7c703b6b1d8dffe1ffae3126b31dd", "x-oauth-basic")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CloseIssue(number string) (*Issue, error) {

	var editReq = IssueRequest{State: "closed"}
	data, err := json.Marshal(editReq)
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s\n", err)
		return nil, err
	}
	req, err := http.NewRequest(
		"PATCH",
		IssueURL+"/"+number,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("d04863358ea7c703b6b1d8dffe1ffae3126b31dd", "x-oauth-basic")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
