package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
    cont := 10
    sum := 0
    
	clearVariable := strings.ReplaceAll(isbn , "-","")
    if len(clearVariable) != 10 {
		return false
	}
    
    
    for i:=0; i<len(clearVariable); i++{

        num:= int(clearVariable[i]- '0' )
        
		if clearVariable[i] < '0' || clearVariable[i] > '9' {
            if i == 9 && clearVariable[i] == 'X'{
                num = 10
            }else{
                return false
            }
            
        }

        sum += num * cont
        cont--
    }

   return sum % 11 == 0
    
}
