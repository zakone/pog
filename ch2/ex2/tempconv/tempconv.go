package tempconv

import "fmt"

type Foot float64
type Metre float64
type Yard float64

func (f Foot) String()  string { return fmt.Sprintf("%gft", f) }
func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (y Yard) String()  string { return fmt.Sprintf("%gyd", y) }
