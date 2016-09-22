// Create: go run issues.go create "create a issue" "test if create a issue success"
// Get All: go run issues.go get all
// Get One: go run issues.go get 1
// Edit: go run issues.go edit 1 "edit a issue"
// Close: go run issues.go close 1
package main

import (
	"./github"
	"fmt"
	"log"
	"os"
)

func main() {
	actionType := os.Args[1]
	if actionType == "create" {
		result, err := github.CreateIssue(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("create issues success: \n")
		fmt.Printf("number: #%d \nuser: %s \ntitle: %s \ncreateAt: %v\n", result.Number, result.User.Login, result.Title, result.CreatedAt)
	} else if actionType == "get" {
		if os.Args[2] == "all" {
			result, err := github.GetIssues()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("get issues success: \n")
			for _, item := range result {
				fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
			}

		} else {
			result, err := github.GetSingleIssue(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("get one issue success: \n")
			fmt.Printf("#%-5d %9.9s %.55s %v\n", result.Number, result.User.Login, result.Title, result.CreatedAt)
		}
	} else if actionType == "edit" {
		result, err := github.EditIssue(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("edit issues success: \n")
		fmt.Printf("#%-5d %9.9s %.55s %v %v\n", result.Number, result.User.Login, result.Title, result.CreatedAt, result.UpdatedAt)
	} else if actionType == "close" {
		result, err := github.CloseIssue(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("close issues success: \n")
		fmt.Printf("#%-5d %9.9s %.55s %.55s %v\n", result.Number, result.User.Login, result.Title, result.State, result.UpdatedAt)
	}

}
