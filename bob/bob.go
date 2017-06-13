package bob

const testVersion = 3

func Hey(s string) string {
	l := string(s[len(s)-1])

	switch l {
	case "?":
		return "Sure."
	case "!":
		return "Whoa, chill out!"
	case "":
		return "Sure."
	default:
		return "Whatever."
	}
}
