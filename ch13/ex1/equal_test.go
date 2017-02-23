package equal

import (
	"testing"
)

func TestFloat64Equal(t *testing.T) {
	var tests = []struct {
		input1 float64
		input2 float64
		want   bool
	}{
		{1.0000000001, 1.0000000002, true},
		{1.0000000009, 1.0000000002, false},
	}
	for _, test := range tests {
		got := Equal(test.input1, test.input2)
		if got != test.want {
			t.Errorf("number %d, %d should be equal, but not", test.input1, test.input2)
		}
	}
}
