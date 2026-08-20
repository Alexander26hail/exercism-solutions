package bottlesong
import "fmt"
import "strings"
var umToLetter = map[int]string{
    1 : "One",
    2 : "Two",
    3 : "Three",
    4 : "Four",
    5 : "Five",
    6 : "Six",
    7 : "Seven",
    8 : "Eight",
    9 : "Nine",
    10: "Ten",
}
func Recite(startBottles, takeDown int) []string {
	Lyric:= []string{}

    for i:=0; i<takeDown; i++{
		

        
        Lyric = append(Lyric , fmt.Sprintf("%s green %s hanging on the wall,",umToLetter[startBottles], bottles(startBottles) ))
        Lyric = append(Lyric , fmt.Sprintf("%s green %s hanging on the wall,",umToLetter[startBottles], bottles(startBottles)  ))
        Lyric = append(Lyric , fmt.Sprintf("And if one green bottle should accidentally fall," ))
        startBottles= startBottles -1
        
        switch {
            
            case startBottles >= 1:
            	Lyric = append(Lyric , fmt.Sprintf("There'll be %s green %s hanging on the wall.",strings.ToLower(umToLetter[startBottles] ), bottles(startBottles) ))
            case startBottles <=0:
            Lyric = append(Lyric , fmt.Sprintf("There'll be no green bottles hanging on the wall."))
        }

        if i < takeDown - 1 {
    		Lyric = append(Lyric, "")
		}
        
    }
    return Lyric
}

func bottles(startBottles int) string {
    bottlestr:=""
    switch {
            case startBottles> 1:
            	bottlestr = "bottles"
            case startBottles == 1:
            	bottlestr = "bottle"
            case startBottles <=0:
            	bottlestr = "bottles"
        }
    return bottlestr
}
