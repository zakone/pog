package palindrome

import "sort"

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i,s.Len()-i-1) || s.Less(s.Len()-i-1,i) {
			return false
		}
	}
	return true
}