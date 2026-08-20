package pangram
import "strings"
func IsPangram(input string) bool {
	mymap:= map[rune]bool{}
    input = strings.ToUpper(input)

    for _ , i := range input{
        
        if i >= 'A' && i <= 'Z' {
            mymap[i]= true
        }
        
        
    }
    return len(mymap) ==26
}
