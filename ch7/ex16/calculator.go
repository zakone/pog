package main

import (
    "fmt"
    "log"
    "net/http"
)

import "./eval"

func parseAndCheck(s string) (eval.Expr, error) {
    if s == "" {
        return nil, fmt.Errorf("empty expression")
    }
    expr, err := eval.Parse(s)
    if err != nil {
        return nil, err
    }
    vars := make(map[eval.Var]bool)
    if err := expr.Check(vars); err != nil {
        return nil, err
    }
    return expr, nil
}

func calculator(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    expr, err := parseAndCheck(r.Form.Get("expr"))
    if err != nil {
        http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "Result: %f", expr.Eval(eval.Env{}))
}

func main() {
    http.HandleFunc("/calculator", calculator)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// localhost:8000/calculator?expr=sin(30)*pow(1.5, 3)
// Result: -3.334607
