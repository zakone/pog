package equal

import (
    "math"
    "reflect"
)

//!+
func equal(x, y reflect.Value) bool {
    if !x.IsValid() || !y.IsValid() {
        return x.IsValid() == y.IsValid()
    }
    if x.Type() != y.Type() {
        return false
    }

    switch x.Kind() {
    //!-
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
        reflect.Int64:
        return x.Int() == y.Int()

    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
        reflect.Uint64, reflect.Uintptr:
        return x.Uint() == y.Uint()

    case reflect.Float32, reflect.Float64:
        return x.Float()-y.Float() <= math.Pow10(-10)

    case reflect.Complex64, reflect.Complex128:
        return x.Complex() == y.Complex()
    default: // chan, func
        return false
    }
}

//!-

//!+comparison
// Equal reports whether x and y are deeply equal.
//!-comparison
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
//!+comparison
func Equal(x, y interface{}) bool {
    return equal(reflect.ValueOf(x), reflect.ValueOf(y))
}
