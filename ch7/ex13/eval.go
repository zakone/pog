package eval

import (
    "bytes"
    "fmt"
    "math"
    "strconv"
    "strings"
    "text/scanner"
)

// An Expr is an arithmetic expression.
type Expr interface {
    // Eval returns the value of this Expr in the environment env.
    Eval(env Env) float64
    // Check reports errors in this Expr and adds its Vars to the set.
    Check(vars map[Var]bool) error
    // Preety Printing
    String() string
}

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
    op rune // one of '+', '-'
    x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
    op   rune // one of '+', '-', '*', '/'
    x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
    fn   string // one of "pow", "sin", "sqrt"
    args []Expr
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
    return env[v]
}

func (l literal) Eval(_ Env) float64 {
    return float64(l)
}

func (u unary) Eval(env Env) float64 {
    switch u.op {
    case '+':
        return +u.x.Eval(env)
    case '-':
        return -u.x.Eval(env)
    }
    panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
    switch b.op {
    case '+':
        return b.x.Eval(env) + b.y.Eval(env)
    case '-':
        return b.x.Eval(env) - b.y.Eval(env)
    case '*':
        return b.x.Eval(env) * b.y.Eval(env)
    case '/':
        return b.x.Eval(env) / b.y.Eval(env)
    }
    panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
    switch c.fn {
    case "pow":
        return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
    case "sin":
        return math.Sin(c.args[0].Eval(env))
    case "sqrt":
        return math.Sqrt(c.args[0].Eval(env))
    }
    panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (v Var) Check(vars map[Var]bool) error {
    vars[v] = true
    return nil
}

func (literal) Check(vars map[Var]bool) error {
    return nil
}

func (u unary) Check(vars map[Var]bool) error {
    if !strings.ContainsRune("+-", u.op) {
        return fmt.Errorf("unexpected unary op %q", u.op)
    }
    return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
    if !strings.ContainsRune("+-*/", b.op) {
        return fmt.Errorf("unexpected binary op %q", b.op)
    }
    if err := b.x.Check(vars); err != nil {
        return err
    }
    return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
    arity, ok := numParams[c.fn]
    if !ok {
        return fmt.Errorf("unknown function %q", c.fn)
    }
    if len(c.args) != arity {
        return fmt.Errorf("call to %s has %d args, want %d",
            c.fn, len(c.args), arity)
    }
    for _, arg := range c.args {
        if err := arg.Check(vars); err != nil {
            return err
        }
    }
    return nil
}

func (v Var) String() string {
    buf := new(bytes.Buffer)
    fmt.Fprintf(buf, "%s", v)
    return buf.String()
    // return fmt.Sprintf("%v", v)
}

func (l literal) String() string {
    return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
    return fmt.Sprintf("(%s%s)", u.op, u.x.String())
}

func (b binary) String() string {
    return fmt.Sprintf("(%s %s %s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
    // return fmt.Sprintf("%s(%s, %s)", c.fn, c.args[0].String(), c.args[1].String())
    buf := new(bytes.Buffer)
    fmt.Fprintf(buf, "%s(", c.fn)
    for i, arg := range c.args {
        if i > 0 {
            buf.WriteString(", ")
        }
        buf.WriteString(arg.String())
    }
    buf.WriteByte(')')
    return buf.String()
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func Format(e Expr) string {
    var buf bytes.Buffer
    write(&buf, e)
    return buf.String()
}

func write(buf *bytes.Buffer, e Expr) {
    switch e := e.(type) {
    case literal:
        fmt.Fprintf(buf, "%g", e)

    case Var:
        fmt.Fprintf(buf, "%s", e)

    case unary:
        fmt.Fprintf(buf, "(%c", e.op)
        write(buf, e.x)
        buf.WriteByte(')')

    case binary:
        buf.WriteByte('(')
        write(buf, e.x)
        fmt.Fprintf(buf, " %c ", e.op)
        write(buf, e.y)
        buf.WriteByte(')')

    case call:
        fmt.Fprintf(buf, "%s(", e.fn)
        for i, arg := range e.args {
            if i > 0 {
                buf.WriteString(", ")
            }
            write(buf, arg)
        }
        buf.WriteByte(')')

    default:
        panic(fmt.Sprintf("unknown Expr: %T", e))
    }
}

type lexer struct {
    scan  scanner.Scanner
    token rune // current lookahead token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
    switch lex.token {
    case scanner.EOF:
        return "end of file"
    case scanner.Ident:
        return fmt.Sprintf("identifier %s", lex.text())
    case scanner.Int, scanner.Float:
        return fmt.Sprintf("number %s", lex.text())
    }
    return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func precedence(op rune) int {
    switch op {
    case '*', '/':
        return 2
    case '+', '-':
        return 1
    }
    return 0
}

// ---- parser ----

// Parse parses the input string as an arithmetic expression.
//
//   expr = num                         a literal number, e.g., 3.14159
//        | id                          a variable name, e.g., x
//        | id '(' expr ',' ... ')'     a function call
//        | '-' expr                    a unary operator (+-)
//        | expr '+' expr               a binary operator (+-*/)
//
func Parse(input string) (_ Expr, err error) {
    defer func() {
        switch x := recover().(type) {
        case nil:
            // no panic
        case lexPanic:
            err = fmt.Errorf("%s", x)
        default:
            // unexpected panic: resume state of panic.
            panic(x)
        }
    }()
    lex := new(lexer)
    lex.scan.Init(strings.NewReader(input))
    lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
    lex.next() // initial lookahead
    e := parseExpr(lex)
    if lex.token != scanner.EOF {
        return nil, fmt.Errorf("unexpected %s", lex.describe())
    }
    return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
    lhs := parseUnary(lex)
    for prec := precedence(lex.token); prec >= prec1; prec-- {
        for precedence(lex.token) == prec {
            op := lex.token
            lex.next() // consume operator
            rhs := parseBinary(lex, prec+1)
            lhs = binary{op, lhs, rhs}
        }
    }
    return lhs
}

// unary = '+' expr | primary
func parseUnary(lex *lexer) Expr {
    if lex.token == '+' || lex.token == '-' {
        op := lex.token
        lex.next() // consume '+' or '-'
        return unary{op, parseUnary(lex)}
    }
    return parsePrimary(lex)
}

// primary = id
//         | id '(' expr ',' ... ',' expr ')'
//         | num
//         | '(' expr ')'
func parsePrimary(lex *lexer) Expr {
    switch lex.token {
    case scanner.Ident:
        id := lex.text()
        lex.next() // consume Ident
        if lex.token != '(' {
            return Var(id)
        }
        lex.next() // consume '('
        var args []Expr
        if lex.token != ')' {
            for {
                args = append(args, parseExpr(lex))
                if lex.token != ',' {
                    break
                }
                lex.next() // consume ','
            }
            if lex.token != ')' {
                msg := fmt.Sprintf("got %q, want ')'", lex.token)
                panic(lexPanic(msg))
            }
        }
        lex.next() // consume ')'
        return call{id, args}

    case scanner.Int, scanner.Float:
        f, err := strconv.ParseFloat(lex.text(), 64)
        if err != nil {
            panic(lexPanic(err.Error()))
        }
        lex.next() // consume number
        return literal(f)

    case '(':
        lex.next() // consume '('
        e := parseExpr(lex)
        if lex.token != ')' {
            msg := fmt.Sprintf("got %s, want ')'", lex.describe())
            panic(lexPanic(msg))
        }
        lex.next() // consume ')'
        return e
    }
    msg := fmt.Sprintf("unexpected %s", lex.describe())
    panic(lexPanic(msg))
}
