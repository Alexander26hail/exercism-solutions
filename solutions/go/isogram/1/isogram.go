package isogram
import "strings"


func IsIsogram(word string) bool {
	maprun := map[rune]bool{}
    word = strings.ToUpper(word)
    

    for _ , i := range word{
        if i == ' ' || i == '-'{
            continue
        }
        
        if maprun[i]{
            return false
        }
        maprun[i] = true
    }
    return true 
}
