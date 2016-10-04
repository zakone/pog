package commaDecimal

import "testing"

func TestCommaDecimal(t *testing.T) {
	var tests = []struct {
		digits   string
		expected string
	}{
		{"", ""},
		{"-1.45", "-1.45"},
		{"-123", "-123"},
		{"-1234", "-1,234"},
		{"1234.56", "1,234.56"},
		{"-12.34567", "-12.34,567"},
	}
	for _, test := range tests {
		if got := commaDecimal(test.digits); got != test.expected {
			t.Errorf("%s IsWrongComma", got)
		}
	}
}
