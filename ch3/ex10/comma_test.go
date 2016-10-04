package comma

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		digits   string
		expected string
	}{
		{"", ""},
		{"1", "1"},
		{"123", "123"},
		{"1234", "1,234"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
	}
	for _, test := range tests {
		if got := comma(test.digits); got != test.expected {
			t.Errorf("%s IsWrongComma", got)
		}
	}
}
