package series
import "fmt"
func All(n int, s string) []string {
    
    mapnumber:= []string{}
    if n > len(s) || n <= 0 {
    	return []string{}  
	}
    for i := 0; i <= len(s) - n; i++{
     	fmt.Print("value i : " , i , "value  maximum , " , len(s) - n ,"   ")
		mapnumber = append(mapnumber,s[i:i+n])

        fmt.Print( " final" , s[i:i+n])
    }
    return mapnumber
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
