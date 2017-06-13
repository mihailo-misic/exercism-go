package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {
	if len(s) == 0 {
		return false
	}

	alph := "abcdefghijklmnopqrstuvwxyz"

	for _, l := range s {
		alph = strings.Replace(alph, strings.ToLower(string(l)), "", -1)
	}

	return len(alph) == 0
}
