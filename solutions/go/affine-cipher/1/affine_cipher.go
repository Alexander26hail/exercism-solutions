package affinecipher
import "errors"
import "strings"
import "unicode"
var abecedario = [26]string{
        "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
        "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}
func Encode(text string, a, b int) (string, error) {
    finalText:=""
    count:=0
   	text= strings.ReplaceAll(strings.ToLower(text)," ", "")
    text= strings.ReplaceAll(text,".", "")
    text= strings.ReplaceAll(text,",", "")
    if GCD(a, 26) != 1 {
        return finalText, errors.New("Error...")
    }
	for i , letter := range text {
        	count++
        	if !unicode.IsLetter(letter) {
        		finalText+=string(letter)
    		}else{
                letterNum:= int(letter - 'a')
                result:= ((a * letterNum + b) %26 +26 ) %26 
                finalText+= abecedario[result]
            	
            }
        	if count == 5 && i+1 != len(text){
                count= 0
                finalText+= " "
            }
        	
            
     
            
    }
    return finalText, nil
}
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func Decode(text string, a, b int) (string, error) {
    original:=""
	text= strings.ReplaceAll(text, " ", "")
    text= strings.ReplaceAll(text, ".","")
    text= strings.ReplaceAll(text, ",","")

    if GCD(a,26)!=1{
        return original, errors.New("Error")
    }
    for _  , i := range text{
		letterToNumber:= int(i - 'a')
        if i>='a' && i <='z'{
            numberLetter:= ((getA(a)) * (letterToNumber - b) %26 + 26) %26
            original+= string(abecedario[numberLetter])
        }else{
            original+=string(i)
        }
        
    }
    return original , nil
    
}

func getA(a int )int{
    for i:=1; i<26; i++{
        if (a*i) %26 ==1{
            return i
        }
    }
    return 0
}
