package secrethandshake
import "strconv"
func Handshake(code uint) []string {
	result := []string{}
    reverse := false
    binary := strconv.FormatUint(uint64(code), 2)
	count := 0
    for i:=len(binary)-1; i>=0; i--{
        count++
        if binary[i]==49{
            switch {
                
                case count==1:
					result = append(result , "wink")
                case count==2:
					result = append(result , "double blink")
                case count==3:
					result = append(result , "close your eyes")
                case count==4:
					result = append(result , "jump")
                case count==5:
					reverse = true
                
            }
        }
        
    }
	if reverse {
    
        resultInvertido := []string{}
        for i:=len(result)-1; i>=0; i--{
            resultInvertido = append(resultInvertido ,result[i] )
        }
        return resultInvertido
        
    }
    
    
    return result
}
