package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvinscale float64

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvinscale) Celsius { return Celsius(k - 273.15) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }


type celsiusFlag struct{ Celsius }


func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvinscale(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
