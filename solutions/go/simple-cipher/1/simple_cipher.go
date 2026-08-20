package cipher
import "unicode"
import "strings"
// Define the shift and vigenere types here.
type shift struct {
    distance int
}

type vigenere struct {
    key string
}
// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return shift{distance:3}
}

func NewShift(distance int) Cipher {
    if distance == 0 || distance > 25 || distance < -25 {
    	return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
  	base := 'a'
    finalResult:=""
    input = strings.ToLower(input)
    for _ , letter := range input{
        if unicode.IsLetter(letter){
        	position := letter - base
            newPosition:= ((position + rune(c.distance) ) %26  + 26 ) %26
            result := base + newPosition
            finalResult+=string(result)
        }
    }
    return finalResult
    
}

func (c shift) Decode(input string) string {
	base := 'a'
    finalResult:=""
    input = strings.ToLower(input)
    for _ , letter := range input{
        if unicode.IsLetter(letter){
            position:= letter - base
            newPosition := ((position - rune(c.distance)) % 26 + 26) % 26
            result:= base + newPosition
            finalResult += string(result)
       }
    }
    return finalResult
}

func NewVigenere(key string) Cipher {
    // Validar que todos sean a-z
    for _, letter := range key {
        if letter < 'a' || letter > 'z' {
            return nil
        }
    }
    

    for _, letter := range key {
        if letter != 'a' {
            return vigenere{key: key}  
        }
    }
    
    // Si llegaste aquí, es solo 'a's
    return nil
}

func (v vigenere) Encode(input string) string {
	base:= 'a'
    finalResult:=""
    input = strings.ToLower(input)
    count:= 0
    for _ , letter := range input{
        if unicode.IsLetter(letter){
            position:= letter - base
            keyShift := rune(v.key[count % len(v.key)]) - 'a'
            newPosition:= ((position + keyShift) % 26 + 26) % 26
            result:= base + newPosition
            finalResult+= string(result)
            count++
        }
    }
    return finalResult
}

func (v vigenere) Decode(input string) string {
	base:='a'
    input= strings.ToLower(input)
    count:=0
    result:=""
    for _ , letter := range input{
        if unicode.IsLetter(letter){
            position:= letter - base
            distance:= rune(v.key[count % len(v.key)]) - 'a'
            newPosition:= ((position - distance) %26 + 26) %26
            result+= string(base + newPosition)
            count++
        }
    }
    return result
}
