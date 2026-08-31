package parallelletterfrequency
import "strings"
import "unicode"
// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
    frequency:= FreqMap{}
    for _ , elem:= range text{
        if unicode.IsLetter(elem) {
        	frequency[elem]++
        }
    }
     
    return frequency
    
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	canal := make(chan FreqMap, len(texts))
	for _ , elem:= range texts {
        
         
        go func(text string){
            canal <- Frequency(text)
        }(strings.ToLower(elem))
        
        
    }
    result := FreqMap{}
    for i:=0; i<len(texts); i++{
        parcial:= <-canal

        for letra , cantidad := range parcial{
            result[letra] += cantidad
        }
    }

    return result
}
