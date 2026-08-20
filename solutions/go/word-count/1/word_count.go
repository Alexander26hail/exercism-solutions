package wordcount
import "strings"
import "fmt"
import "unicode"
type Frequency map[string]int

func WordCount(phrase string) Frequency {
    
    clened := strings.Map(func(r rune) rune {
    if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || r == '\'' {
        return r
    }
    	return ' '  
	}, phrase)
    
    words := strings.Fields(clened)
    
    
    frequency:= Frequency{}
    fmt.Printf("%q",words )
	for _ , spli := range words {
        
        spli = strings.ToLower(spli)
        spli = strings.Trim(spli, "'")
        if spli == ""{
            continue
        }
        frequency[spli]+= 1
    }
    
	return frequency
}


