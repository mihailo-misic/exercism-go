package secret

import "fmt"

const testVersion = 2

func Handshake(i uint) []string {
	r := make([]string, 0, 4)
	if i < 1 {
		return r
	}

	// Turn the uint into binary.
	b := fmt.Sprintf("%05b", i)

	if string(b[4]) == "1" { // Check the first digit.
		r = append(r, "wink")
	}
	if string(b[3]) == "1" { // Check the second digit.
		r = append(r, "double blink")
	}
	if string(b[2]) == "1" { // Check the third digit.
		r = append(r, "close your eyes")
	}
	if string(b[1]) == "1" { // Check the fourth digit.
		r = append(r, "jump")
	}

	// Reverse the array if the over 10000
	if i > 15 {
		return reverseArray(r)
	}

	return r
}

// reverseArray function reverses the given array of strings
func reverseArray(array []string) []string {
	if len(array) == 0 {
		return array
	}
	return append(reverseArray(array[1:]), array[0])
}
