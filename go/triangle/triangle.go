package triangle

import (
	"math"
)

// Kind is the type of triangle
type Kind string

// Possible types of triangles
const (
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// KindFromSides check the sides of a triangle and determines the type
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else {
		k = Sca
	}

	// Check triangle inequality
	if a+b < c || a+c < b || b+c < a {
		k = NaT
	}

	// Check if any side is NaN
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
	}

	// Check if any side is Inf
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	}

	return k
}
