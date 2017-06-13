package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Hey(s string) string {
	r := ""
	s = strings.TrimSpace(s)
	yelling, _ := regexp.MatchString(`([A-Z]{4}|[A-Z]!)`, s)

	if len(s) == 0 {
		return "Fine. Be that way!"
	} else if yelling {
		return "Whoa, chill out!"
	}

	l := string(s[len(s)-1])

	if l == "?" {
		r = "Sure."
	} else if l == "\t" {
		r = "Fine. Be that way!"
	} else {
		r = "Whatever."
	}

	return r
}
