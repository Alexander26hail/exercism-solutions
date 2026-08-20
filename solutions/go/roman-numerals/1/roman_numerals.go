package romannumerals

import "errors"
func ToRomanNumeral(input int) (string, error) {
    roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    arabicletter:=""
	if input >0 && input <4000{
        for i , x := range arabic{
            for  input >= x{
                input -= x
                arabicletter += roman[i]
            }
        }
        return arabicletter , nil
    }
    return arabicletter, errors.New("Error")
}
