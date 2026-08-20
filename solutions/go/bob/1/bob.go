// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob
import "unicode"
import "strings"
// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch  {
        case  strings.HasSuffix(remark, "?") &&  !isYelling(remark):
        	return "Sure."
        case !strings.HasSuffix(remark, "?") &&   isYelling(remark):
        	return "Whoa, chill out!"

        case strings.HasSuffix(remark, "?") &&   isYelling(remark):
			return "Calm down, I know what I'm doing!"

        case  remark == "" :
        	return "Fine. Be that way!"
        default: 
        	return "Whatever." 
    }
    
}

func isYelling(s string) bool {
    hasLetter := false
    for _, r := range s {
        if unicode.IsLetter(r) {
            hasLetter = true
            if !unicode.IsUpper(r) {
                return false  // encontró una minúscula
            }
        }
    }
    return hasLetter  // grita si tiene letras Y todas son mayúsculas
}
