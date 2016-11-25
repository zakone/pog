package palindrome

import "testing"

type Palindrome struct {
	data []byte
}

func (x Palindrome) Len() int           { return len(x.data) }
func (x Palindrome) Less(i, j int) bool { return x.data[i] < x.data[j] }
func (x Palindrome) Swap(i, j int)      { x.data[i], x.data[j] = x.data[j], x.data[i] }

func TestIsPalindrome(t *testing.T) {
	var s Palindrome
	var tests = []struct {
		input []byte
		want  bool
	}{
		{[]byte(""), true},
		{[]byte("abcba"), true},
		{[]byte("not a palindrome"), false},
	}
	for _, test := range tests {
		s.data = test.input
		if got := IsPalindrome(s); got != test.want {
			t.Errorf("%s IsPalindrome: %t, Should be %t", test.input, got, test.want)
		}
	}
}
