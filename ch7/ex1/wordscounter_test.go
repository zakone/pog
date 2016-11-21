package wordscounter

import "testing"
import "fmt"

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte("hello world golang practice"), 4},
		{[]byte(""), 0},
		{[]byte("golanggolanggolang"), 1},
	}
	for _, test := range tests {
		var c WordCounter
		if got, _ := c.Write(test.input); got != test.want {
			t.Errorf("%s\t Wrong Words Count %d, Should be %d", test.input, got, test.want)
		}
		fmt.Println(c)
	}
}

func TestLineCounter(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte("hello world golang practice\n"), 1},
		{[]byte(""), 0},
		{[]byte("golang\ngolang\ngolang"), 3},
	}
	for _, test := range tests {
		var c LineCounter
		if got, _ := c.Write(test.input); got != test.want {
			t.Errorf("%s\t Wrong Line Count %d, Should be %d", test.input, got, test.want)
		}
		fmt.Println(c)
	}
}
