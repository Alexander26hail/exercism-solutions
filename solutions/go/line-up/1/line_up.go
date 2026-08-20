package lineup
import "fmt"

func Format(name string, number int) string {
    ordinal:="th"
	FistDigit:= number %10
    SecondNumber:= number %100
    if SecondNumber!= 11 && SecondNumber!= 12 && SecondNumber!= 13 {
        switch {
            case FistDigit == 1 :
            	ordinal = "st"
            case FistDigit == 2 :
            	ordinal = "nd"
    		case FistDigit == 3 :
            	ordinal = "rd"
            default :
            	ordinal = "th"
        }
    }
    
	if number >= 1 && number <= 999{
        return fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!",name , number , ordinal)
    }
    return ""
}
