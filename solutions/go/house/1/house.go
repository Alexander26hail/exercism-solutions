package house
import "strings"
var list1 = map[int]string{
    1: "This is the house that Jack built.",
    2:"malt",
    3:"rat",
    4:"cat",
    5:"dog",
    6:"cow with the crumpled horn",
    7:"maiden all forlorn",
    8:"man all tattered and torn",
    9:"priest all shaven and shorn",
    10:"rooster that crowed in the morn",
    11:"farmer sowing his corn",
    12:"horse and the hound and the horn",

    
}

var list2 = map[int]string{
    1: "This is the house that Jack built.",
    2:"that ate the malt",
    3:"that killed the rat",
    4:"that worried the cat",
    5:"that tossed the dog",
    6:"that milked the cow with the crumpled horn",
    7:"that kissed the maiden all forlorn",
    8:"that married the man all tattered and torn",
    9:"that woke the priest all shaven and shorn",
    10:"that kept the rooster that crowed in the morn",
    11:"that belonged to the farmer sowing his corn",


    
}
func Verse(v int) string {
    verse:=[]string{}
    finalVerse:="that lay in the house that Jack built."
	if v<= 0{
        return ""
    }
    switch {
        case v == 1:
        	verse = append(verse, list1[v])

        default :
        	verse = append(verse , "This is the " + list1[v] )
        	for i:=v-1; i >= 2; i--{ 
                verse = append(verse ,  list2[i] )
            }
        	verse = append(verse , finalVerse)
    }

    return strings.Join(verse , "\n")
}

func Song() string {
	verse:=[]string{}

    for i:=1; i<=12; i++{
        verse = append(verse , Verse(i))
    }
    return strings.Join(verse, "\n\n")
}
