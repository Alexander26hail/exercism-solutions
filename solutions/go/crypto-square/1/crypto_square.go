package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	var builder strings.Builder
	for _, char := range pt {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(unicode.ToLower(char))
		}
	}
	pt = builder.String()

	if pt == "" {
		return pt
	}

	r := int(math.Sqrt(float64(len(pt))))
	var c int


	for {
		c = (len(pt) + r - 1) / r
		if c >= r && c-r <= 1 {
			break
		}
		r++
	}

	array := []string{}
	result := []string{}

	for i := 0; i < len(pt); i = i + c {
		
		end := i + c
		if end > len(pt) {
			end = len(pt)
		}
		extract := pt[i:end]

		
		if len(extract) < c {
			extract += strings.Repeat(" ", c-len(extract))
		}
		array = append(array, extract)
	}

	for col := 0; col < c; col++ {
		var chunks string

		for row := 0; row < r; row++ {
			if row < len(array) {
				chunks += string(array[row][col])
			} else {
				chunks += " "
			}
		}
		
		result = append(result, chunks)
	}

	return strings.Join(result, " ")
}