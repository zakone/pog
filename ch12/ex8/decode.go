package sexpr

import (
    "bytes"
    "fmt"
    "io"
    "reflect"
    "strconv"
    "text/scanner"
)

//!+Unmarshal
// Unmarshal parses S-expression data and populates the variable
// whose address is in the non-nil pointer out.
func unmarshal(data []byte, out interface{}) (err error) {
    dec := &Decoder{scan: scanner.Scanner{Mode: scanner.GoTokens}}
    dec.scan.Init(bytes.NewReader(data))
    dec.next() // get the first token
    defer func() {
        // NOTE: this is not an example of ideal error handling.
        if x := recover(); x != nil {
            err = fmt.Errorf("error at %s: %v", dec.scan.Position, x)
        }
    }()
    read(dec, reflect.ValueOf(out).Elem())
    return nil
}

//!-Unmarshal

//!+Decoder
// type Decoder struct {
//     scan  scanner.Scanner
//     token rune // the current token
// }

type Decoder struct {
    r     io.Reader
    buf   []byte
    scanp int
    scan  scanner.Scanner
    token rune
}

func NewDecoder(r io.Reader) *Decoder {
    return &Decoder{r: r}
}

func (dec *Decoder) Decode(v interface{}) error {
    // Read whole value into buffer.
    dec.readValue()
    err := unmarshal(dec.buf, v)
    dec.scanp += len(dec.buf)
    return err
}

func (dec *Decoder) readValue() {
    buf := new(bytes.Buffer)
    buf.ReadFrom(dec.r)
    b := buf.Bytes()
    if len(dec.buf) == 0 { //最初bufにデータない時
        dec.buf = b
        dec.scanp = 0

    }
    if len(dec.buf)-dec.scanp >= len(b) { //bufまだ余ってる時
        dec.buf = append(dec.buf, b...)

    }
    if len(dec.buf)-dec.scanp < len(b) { //buf余ってる分足りない時
        dec.refill(b)

    }

}

func (dec *Decoder) refill(b []byte) {

    if dec.scanp > 0 {
        n := copy(dec.buf, dec.buf[dec.scanp:])
        dec.buf = dec.buf[:n]
        dec.scanp = 0
    }
    if cap(dec.buf)-len(dec.buf) < len(b) {
        newBuf := make([]byte, len(dec.buf), 2*(cap(dec.buf)+len(b)))
        copy(newBuf, dec.buf)
        dec.buf = newBuf
    }
    dec.buf = append(dec.buf, b...)

}

func (dec *Decoder) next()        { dec.token = dec.scan.Scan() }
func (dec *Decoder) text() string { return dec.scan.TokenText() }

func (dec *Decoder) consume(want rune) {
    if dec.token != want { // NOTE: Not an example of good error handling.
        panic(fmt.Sprintf("got %q, want %q", dec.text(), want))
    }
    dec.next()
}

//!-Decoder

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
// - that v is always a variable of the appropriate type for the
//   S-expression value.  For example, v must not be a boolean,
//   interface, channel, or function, and if v is an array, the input
//   must have the correct number of elements.
// - that v in the top-level call to read has the zero value of its
//   type and doesn't need clearing.
// - that if v is a numeric variable, it is a signed integer.

//!+read
func read(dec *Decoder, v reflect.Value) {
    switch dec.token {
    case scanner.Ident:
        // The only valid identifiers are
        // "nil" and struct field names.
        if dec.text() == "nil" {
            v.Set(reflect.Zero(v.Type()))
            dec.next()
            return
        }
    case scanner.String:
        s, _ := strconv.Unquote(dec.text()) // NOTE: ignoring errors
        v.SetString(s)
        dec.next()
        return
    case scanner.Int:
        i, _ := strconv.Atoi(dec.text()) // NOTE: ignoring errors
        v.SetInt(int64(i))
        dec.next()
        return
    case '(':
        dec.next()
        readList(dec, v)
        dec.next() // consume ')'
        return
    }
    panic(fmt.Sprintf("unexpected token %q", dec.text()))
}

//!-read

//!+readlist
func readList(dec *Decoder, v reflect.Value) {
    switch v.Kind() {
    case reflect.Array: // (item ...)
        for i := 0; !endList(dec); i++ {
            read(dec, v.Index(i))
        }

    case reflect.Slice: // (item ...)
        for !endList(dec) {
            item := reflect.New(v.Type().Elem()).Elem()
            read(dec, item)
            v.Set(reflect.Append(v, item))
        }

    case reflect.Struct: // ((name value) ...)
        for !endList(dec) {
            dec.consume('(')
            if dec.token != scanner.Ident {
                panic(fmt.Sprintf("got token %q, want field name", dec.text()))
            }
            name := dec.text()
            dec.next()
            read(dec, v.FieldByName(name))
            dec.consume(')')
        }

    case reflect.Map: // ((key value) ...)
        v.Set(reflect.MakeMap(v.Type()))
        for !endList(dec) {
            dec.consume('(')
            key := reflect.New(v.Type().Key()).Elem()
            read(dec, key)
            value := reflect.New(v.Type().Elem()).Elem()
            read(dec, value)
            v.SetMapIndex(key, value)
            dec.consume(')')
        }

    default:
        panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
    }
}

func endList(dec *Decoder) bool {
    switch dec.token {
    case scanner.EOF:
        panic("end of file")
    case ')':
        return true
    }
    return false
}

//!-readlist
