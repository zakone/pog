package splitTest

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"a:b:c", 3},
		{"4:5:6", 3},
		{"", 1},
		{"1111", 1},
	}
	for _, test := range tests {
		words := strings.Split(test.input, ":")
		got := len(words)
		if got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.input, ":", got, test.want)
		}
	}

}
