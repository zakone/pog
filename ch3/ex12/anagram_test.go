package anagram

import "testing"

func TestAnagram(t *testing.T) {
	var tests = []struct {
		para1    string
		para2    string
		expected bool
	}{
		{"anagrams", "ARS MAGNA", true},
	}
	for _, test := range tests {
		if got := Anagram(test.para1, test.para2); got != test.expected {
			t.Errorf("%s IsWrongComma", got)
		}
	}
}
