// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
    "fmt"
    "math"
    "testing"
)

//!+Eval
func TestEval(t *testing.T) {
    tests := []struct {
        expr string
        env  Env
        want string
    }{
        {"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
        {"pow(x,3)+pow(y,3)", Env{"x": 12, "y": 1}, "1729"},
        {"pow(x,3)+pow(y,3)", Env{"x": 9, "y": 10}, "1729"},
        {"5/9*(F-32)", Env{"F": -40}, "-40"},
        {"5/9*(F-32)", Env{"F": 32}, "0"},
        {"5/9*(F-32)", Env{"F": 212}, "100"},
        {"-1+-x", Env{"x": 1}, "-2"},
        {"-1-x", Env{"x": 1}, "-2"},
    }
    for _, test := range tests {
        // Print expr only when it changes.
        fmt.Printf("\n%s\n", test.expr)
        expr, err := Parse(test.expr)
        if err != nil {
            t.Error(err) // parse error
            continue
        }

        fmt.Println(expr)
        // s := fmt.Sprintf("%s", expr)
        // exprUpdated, err := Parse(s)
        if err != nil {
            t.Error(err) // parse error
            continue
        }
        got := fmt.Sprintf("%.6g", expr.Eval(test.env))
        fmt.Printf("\t%v => %s\n", test.env, got)
        if got != test.want {
            t.Errorf("%s.Eval() in %v = %q, want %q\n",
                expr, test.env, got, test.want)
        }
    }
}
