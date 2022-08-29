//Package triangle gives us tools to work with triangles
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "Nat" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	//var k Kind
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a != b && a != c && b != c {
		if legality(a, b, c) {
			return Sca
		}
	}
	if (a == b && a != c) || (a == c && a != b) || (c == b && c != a) {
		if legality(a, b, c) {
			return Iso
		}
	}
	return NaT
}

func legality(a, b, c float64) bool {
	if (a+b >= c) || (a+c >= b) || (b+c >= a) {
		return true
	}
	return false
}
