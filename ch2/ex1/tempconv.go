package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvinscale float64

const (
	AbsolutZeroC Celsius     = -273.15
	FreezingC    Celsius     = 0
	BoilingC     Celsius     = 100
	ZeroK        Kelvinscale = 0
	FreezingK    Kelvinscale = 273.15
	BoilingK     Kelvinscale = 373.15
)

func (c Celsius) String() string     { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string  { return fmt.Sprintf("%g°F", f) }
func (k Kelvinscale) String() string { return fmt.Sprintf("%g°K", k) }
