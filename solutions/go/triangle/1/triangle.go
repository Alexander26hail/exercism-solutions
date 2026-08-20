// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
// type Kind
type Kind int 

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota
	Equ Kind = iota
	Iso Kind = iota
	Sca Kind = iota
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	
    switch {
        case a <= 0 || b <= 0 || c <= 0 || (a + b) <= c || (b + c) <= a || (a + c) <= b:
            return NaT
        case a == b && b == c :
        	return Equ
        case a != b && b != c && a != c:
        	return Sca
        default:
        	return Iso
        
    }
}
