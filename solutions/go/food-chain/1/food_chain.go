package foodchain
import "fmt"
import "strings"
var tabla1 = map[int]string{
    1: "fly",
    2: "spider",
    3: "bird",
    4: "cat",
    5: "dog",
    6: "goat",
    7: "cow",
    8: "horse",
}

var tabla2 = map[int]string{
    1: "",
    2: "It wriggled and jiggled and tickled inside her.",
    3: "How absurd to swallow a bird!",
    4: "Imagine that, to swallow a cat!",
    5: "What a hog, to swallow a dog!",
    6: "Just opened her throat and swallowed a goat!",
    7: "I don't know how she swallowed a cow!",
    8: "",
}
func Verse(v int) string {
    verse :=[]string{}
	switch {
        case v==1:
			verse = append(verse ,fmt.Sprintf("I know an old lady who swallowed a fly.\nI don't know why she swallowed the fly. Perhaps she'll die."))
        case v == 8:
			verse = append(verse ,fmt.Sprintf("I know an old lady who swallowed a horse.\nShe's dead, of course!"))
        default:
        	verse = append(verse , fmt.Sprintf("I know an old lady who swallowed a %s.",tabla1[v]))

        verse = append(verse , string(tabla2[v]))
        
        	for i := v; i >= 2; i-- {
                animalActual := tabla1[i]
    			animalAnterior := tabla1[i-1]

                if animalAnterior =="spider"{
                    verse = append(verse , fmt.Sprintf("She swallowed the %s to catch the spider that wriggled and jiggled and tickled inside her.",animalActual ))
                    
                }else{
                    verse = append(verse , fmt.Sprintf("She swallowed the %s to catch the %s.",animalActual, animalAnterior ))
                }
                
                
            }
        verse = append(verse , "I don't know why she swallowed the fly. Perhaps she'll die.")
    }
    return strings.Join(verse, "\n")
}

func Verses(start, end int) string {
	verse :=[]string{}

    for i := start; i <= end; i++ {
        
        verse = append(verse, Verse(i))
    }
    
    return strings.Join(verse, "\n\n")
}

func Song() string {
	return Verses(1,8)
}
