package parsinglogfiles
import "regexp"
import "fmt"
func IsValidLine(text string) bool {
    if text=="" || len(text) <=5 {
        return false
    }

   
    re := regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[ERR\]|^\[INF\]|^\[WRN\]|^\[FTL\]`)
    return re.MatchString(text)
   
	
    
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<[-=*~]*\>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    re := regexp.MustCompile(`"(?i).*password.*"`)
    count:=0
    for _ , i := range lines{
        if re.MatchString(i){
            count++
        }
        
    }
    return count
	
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(?i)end-of-line[0-9]+`)
    fmt.Sprintf("print %s",re)
    return re.ReplaceAllString(text, "") 
}

func TagWithUserName(lines []string) []string {
    result := []string{}
	re := regexp.MustCompile(`User\s+(\S+)`)
    for _ , i:= range lines{
        value:= re.FindStringSubmatch(i)
        if value !=nil{
            username := value[1]
            text:= fmt.Sprintf("[USR] %s %s",username , i)
        	result = append(result,text)
        }else{
            result = append(result,i)
        }
        
    }
    return result
    
    
}
