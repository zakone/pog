package main

import (
	"./github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	t := time.Now()
	durMonth := make([]string, 0, len(result.Items))
	durInYear := make([]string, 0, len(result.Items))
	durOverYear := make([]string, 0, len(result.Items))
	for _, item := range result.Items {
		duration := t.Sub(item.CreatedAt)
		hours := int(duration.Hours())
		month := float64(hours) / float64(24*30)
		if month < 1.0 {
			durMonth = append(durMonth, fmt.Sprintf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt))
		} else if month > 1.0 && month < 12.0 {
			durInYear = append(durInYear, fmt.Sprintf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt))
		} else {
			durOverYear = append(durOverYear, fmt.Sprintf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt))
		}
	}
	fmt.Println("Issue in month:")
	for _, val := range durMonth {
		fmt.Print(val)
	}
	fmt.Println("Issue in year:")
	for _, val := range durInYear {
		fmt.Print(val)
	}
	fmt.Println("Issue over year:")
	for _, val := range durOverYear {
		fmt.Print(val)
	}
}
