package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Abbreviate(s string) string {
	r := ""
	re := regexp.MustCompile(`\w+`) // I'm interested in words only

	split := re.FindAllString(s, -1) // Split the words up

	for _, v := range split {
		r += strings.ToUpper(string(v[0])) // Append the UC first letters to result
	}

	return r
}
