package display2

import (
	"testing"
)

func TestDisplay(t *testing.T) {
	type Cycle struct { Value int; Tail *Cycle }
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c, 10)

	// Output
	// Display c (main.Cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail).Value = 42
	// (*(*(*c.Tail).Tail).Tail).Value = 42
	// (*(*(*(*c.Tail).Tail).Tail).Tail).Value = 42
	// (*(*(*(*(*c.Tail).Tail).Tail).Tail).Tail).Value = 42
	// (*(*(*(*(*(*c.Tail).Tail).Tail).Tail).Tail).Tail).Value = 42
	// (*(*(*(*(*(*(*c.Tail).Tail).Tail).Tail).Tail).Tail).Tail).Value = 42
	// (*(*(*(*(*(*(*(*c.Tail).Tail).Tail).Tail).Tail).Tail).Tail).Tail).Value = 42
	// (*(*(*(*(*(*(*(*(*c.Tail).Tail).Tail).Tail).Tail).Tail).Tail).Tail).Tail).Value = 42

}
