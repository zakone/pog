package main

import (
    "bytes"
    "fmt"
    "reflect"
    "strconv"
)

func main() {
    //!+movie
    type Movie struct {
        Title, Subtitle string
        Year            int
        Color           bool
        Actor           map[string]string
        Oscars          []string
        Sequel          *string
    }
    //!-movie
    //!+strangelove
    strangelove := Movie{
        Title:    "Dr. Strangelove",
        Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
        Year:     1964,
        Color:    false,
        Actor: map[string]string{
            "Dr. Strangelove":            "Peter Sellers",
            "Grp. Capt. Lionel Mandrake": "Peter Sellers",
            "Pres. Merkin Muffley":       "Peter Sellers",
            "Gen. Buck Turgidson":        "George C. Scott",
            "Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
            `Maj. T.J. "King" Kong`:      "Slim Pickens",
        },

        Oscars: []string{
            "Best Actor (Nomin.)",
            "Best Adapted Screenplay (Nomin.)",
            "Best Director (Nomin.)",
            "Best Picture (Nomin.)",
        },
    }
    //!-strangelove
    output, _ := Marshal(strangelove)
    fmt.Printf("%s", output)

    //Output
    // ((Title "Dr. Strangelove")
    //  (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
    //  (Year 1964)
    //  (Color nil)
    //  (Actor (("Gen. Buck Turgidson" "George C. Scott")
    //     ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
    //     ("Maj. T.J. \"King\" Kong" "Slim Pickens")
    //     ("Dr. Strangelove" "Peter Sellers")
    //     ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
    //     ("Pres. Merkin Muffley" "Peter Sellers")))
    //  (Oscars ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)" "Best Director (Nomin.)" "Best Picture (Nomin.)"))
    //  (Sequel nil))
}

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
    var buf bytes.Buffer
    if err := encode(&buf, reflect.ValueOf(v)); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value) error {
    switch v.Kind() {
    case reflect.Invalid:
        buf.WriteString("nil")

    case reflect.Int, reflect.Int8, reflect.Int16,
        reflect.Int32, reflect.Int64:
        fmt.Fprintf(buf, "%d", v.Int())

    case reflect.Uint, reflect.Uint8, reflect.Uint16,
        reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        fmt.Fprintf(buf, "%d", v.Uint())

    case reflect.String:
        fmt.Fprintf(buf, "%q", v.String())

    case reflect.Ptr:
        return encode(buf, v.Elem())

    case reflect.Array, reflect.Slice: // (value ...)
        buf.WriteByte('(')
        for i := 0; i < v.Len(); i++ {
            if i > 0 {
                buf.WriteByte(' ')
            }
            if err := encode(buf, v.Index(i)); err != nil {
                return err
            }
        }
        buf.WriteByte(')')

    case reflect.Struct: // ((name value) ...)
        buf.WriteByte('(')
        for i := 0; i < v.NumField(); i++ {
            if i > 0 {
                buf.WriteByte(' ')
            }
            fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
            if err := encode(buf, v.Field(i)); err != nil {
                return err
            }
            buf.WriteByte(')')
            if i < v.NumField()-1 {
                buf.WriteString("\n")
            }

        }
        buf.WriteByte(')')
        buf.WriteString("\n")

    case reflect.Map: // ((key value) ...)
        buf.WriteByte('(')
        for i, key := range v.MapKeys() {
            if i > 0 {
                buf.WriteByte('\t')
            }
            buf.WriteByte('(')
            if err := encode(buf, key); err != nil {
                return err
            }
            buf.WriteByte(' ')
            if err := encode(buf, v.MapIndex(key)); err != nil {
                return err
            }
            buf.WriteByte(')')
            if i < len(v.MapKeys())-1 {
                buf.WriteString("\n")
            }
        }
        buf.WriteByte(')')

    case reflect.Bool:
        if v.Bool() == true {
            buf.WriteString("t")
        } else {
            buf.WriteString("nil")
        }

    case reflect.Float32, reflect.Float64:
        buf.WriteString(strconv.FormatFloat(v.Float(), 'g', -1, 64))

    case reflect.Complex64, reflect.Complex128: // #C(1.0 2.0)
        buf.WriteString("#C(")
        buf.WriteString(strconv.FormatFloat(real(v.Complex()), 'g', -1, 64))
        buf.WriteByte(' ')
        buf.WriteString(strconv.FormatFloat(imag(v.Complex()), 'g', -1, 64))
        buf.WriteByte(')')

    case reflect.Interface: //("[]int" (1 2 3))
        if v.IsNil() {
            buf.WriteString("nil")
        } else {
            buf.WriteByte('(')
            buf.WriteString(v.Type().String())
            buf.WriteByte(' ')
            encode(buf, v.Elem())
            buf.WriteByte(')')
            buf.WriteString("\n")
        }
    default: // chan, func
        return fmt.Errorf("unsupported type: %s", v.Type())
    }
    return nil
}

//!-encode
