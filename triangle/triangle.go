package triangle

import "math"

const testVersion = 3

type Kind int

var NaT Kind = -1 // not a triangle
var Sca Kind = 0  // 0 same
var Iso Kind = 1  // 2 same
var Equ Kind = 3  // 3 same

func KindFromSides(a, b, c float64) Kind {
	var r Kind

	// Check if not a triangle
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) { // Some side is NaN
		return NaT
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) { // Some side is NaN
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 { // No side can be <= 0
		return NaT
	}
	if a+b < c || b+c < a || c+a < b { // Fails triangle inequity
		return NaT
	}

	// Count number of same sides
	if a == b {
		r++
	}
	if b == c {
		r++
	}
	if c == a {
		r++
	}

	return r
}
