// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
    "fmt"
    "testing"
)

//!+Eval
func TestOperandEval(t *testing.T) {
    tests := []struct {
        expr string
        env  Env
        want string
    }{
        {"min(x, y)", Env{"x": 1, "y": 2}, "1"},
        {"max(x, y)", Env{"x": 1, "y": 2}, "2"},
    }
    var prevExpr string
    for _, test := range tests {
        // Print expr only when it changes.
        if test.expr != prevExpr {
            fmt.Printf("\n%s\n", test.expr)
            prevExpr = test.expr
        }
        expr, err := Parse(test.expr)
        if err != nil {
            t.Error(err) // parse error
            continue
        }
        got := fmt.Sprintf("%.6g", expr.Eval(test.env))
        fmt.Printf("\t%v => %s\n", test.env, got)
        if got != test.want {
            t.Errorf("%s.Eval() in %v = %q, want %q\n",
                test.expr, test.env, got, test.want)
        }
    }
}

func TestErrors(t *testing.T) {
    for _, test := range []struct{ expr, wantErr string }{
        {"log(10)", `unknown function "log"`},
        {"min(1)", "call to min has 1 args, want 2"},
        {"max(1)", "call to max has 1 args, want 2"},
    } {
        expr, err := Parse(test.expr)
        if err == nil {
            vars := make(map[Var]bool)
            err = expr.Check(vars)
            if err == nil {
                t.Errorf("unexpected success: %s", test.expr)
                continue
            }
        }
        fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
        if err.Error() != test.wantErr {
            t.Errorf("got error %s, want %s", err, test.wantErr)
        }
    }
}
