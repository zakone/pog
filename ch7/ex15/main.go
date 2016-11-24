package main

import (
	"./eval"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Expression: ")
		express, _ := reader.ReadString('\n')
		express = strings.TrimRight(express, "\n")
		if express == "" {
			fmt.Println("Empty Expression")
		}
		expr, err := eval.Parse(express)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unknown Expression:", err)
			continue
		}
		vars := make(map[eval.Var]bool)
		err = expr.Check(vars)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Expression Format Error:", err)
			continue
		}
		env := make(eval.Env, 0)
		count := 0
		for k, ok := range vars {
			if !ok {
				continue
			}
			fmt.Printf("Please Input Value For %s\n", k)
			input, _ := reader.ReadString('\n')
			input = strings.TrimRight(input, "\n")
			if input == "" {
				fmt.Println("Empty Input")
				continue
			}
			inputFloat64, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "strconv Error:", err)
				continue

			}
			env[k] = inputFloat64
			count++
		}
		if len(env) != count {
			continue
		}
		fmt.Printf("Result: %.6g\n", expr.Eval(env))
	}

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
