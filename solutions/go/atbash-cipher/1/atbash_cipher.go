package atbashcipher
import "strings"
import "unicode"
var PlanToCipher= map[rune]rune{}

func convert(){
    plain := "abcdefghijklmnopqrstuvwxyz"
    cipher := "zyxwvutsrqponmlkjihgfedcba"

    for i , p := range plain{
        PlanToCipher[p] = rune(cipher[i])
    }
}
func Atbash(s string) string {
    convert()
    text:= ""
    s = strings.Map(func(r rune) rune {
    if unicode.IsLetter(r) || unicode.IsDigit(r) {
        return r
    }
    	return -1  // elimina
	}, s)
    count :=0
    s = strings.ToLower(s)
	for _ , i := range s{
		
        value , exist := PlanToCipher[i]
		if exist { 
            
			text += string(value)
        }else{
            text += string(i)
        }
        count++
        if count == 5 {
                count= 0 
                text += " "
        }
        
        
        
       
    } 
    return strings.TrimSpace(text) 
}
