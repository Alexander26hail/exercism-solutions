package logs
import "unicode/utf8"
// Application identifies the application emitting the given log.
func Application(log string) string {
    if log !=""{
        for _ , logs := range log{
            
            if logs == '❗'{
                return "recommendation"
            }
            if logs == '🔍'{
                return "search"
            }
            if logs == '☀'{
                return "weather"
            }
            
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    text:= ""
	for _ , runes := range log{
        if runes == oldRune {
            runes = newRune
        }
        text = text + string(runes)

        
    }
    return text
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    return utf8.RuneCountInString(log) <= limit

    return false 
}
