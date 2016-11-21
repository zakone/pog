package writercounter

import "testing"
import "bytes"

func TestWriteCounter(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int64
	}{
		{[]byte("hello golang"), 12},
		{[]byte(""), 0},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		counter, ptr := CountingWriter(&buf)
		counter.Write(test.input)
		if *ptr != test.want {
			t.Errorf("Wrong Writer Count %d, Should be %d", test.input, *ptr, test.want)
		}
	}
}
