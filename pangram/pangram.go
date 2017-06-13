package pangram

import "strings"

const testVersion = 1

func IsPangram(s string) bool {
	s = strings.ToLower(s)
	alph := "abcdefghijklmnopqrstuvwxyz"

	for _, l := range alph {
		if !strings.ContainsAny(s, string(l)) {
			return false
		}
	}

	return true
}
