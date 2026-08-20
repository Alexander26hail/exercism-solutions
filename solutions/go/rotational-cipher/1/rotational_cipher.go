package rotationalcipher

import "strings"
import "unicode"

const alphabet = "abcdefghijklmnopqrstuvwxyz"
func RotationalCipher(plain string, shiftKey int) string {
    text:=""
    for _, i := range plain{
        if unicode.IsLetter(i){
            value:= strings.Index(alphabet, strings.ToLower(string(i)) )
            shiftKey2:= (shiftKey + value)% 26 

            
            if unicode.IsUpper(i){
                text+= strings.ToUpper(string(alphabet[shiftKey2]))
            }else{
                text+= string(alphabet[shiftKey2])
        	}
 
        }else{
             text += string(i)
        }
       
           
    }
    
	return text
    
}
