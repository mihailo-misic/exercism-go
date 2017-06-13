package triangle

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {

}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.
var NaT = 0 // not a triangle
var Equ = 1 // equilateral
var Iso = 2 // isosceles
var Sca = 3 // scalene
