package diamond
import "fmt"
import "strings"
import "errors"
func Gen(char byte) (string, error) {
    alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    finalText:= []string{}
	

    if char < 'A' || char > 'Z'{
        return "" , errors.New("error")
    }

        position:= int(char)-65
       
        for i:=0; i<=position; i++{
            space:= strings.Repeat(" ", position - i)
            if (2*i)-1 >0 {
                internos:= strings.Repeat(" ", 2*i-1)
                finalText = append(finalText , fmt.Sprintf("%s%s%s%s%s", space, string(alphabet[i]),internos,string(alphabet[i]), space))
            }else{
                finalText = append(finalText, fmt.Sprintf("%s%s%s", space, string(alphabet[i]), space))
            }
   
        }

        for x:=position-1; x>=0; x--{
                space:= strings.Repeat(" ", position - x)
            if (2*x)-1 >0 {
                internos:= strings.Repeat(" ", 2*x-1)
                finalText = append(finalText , fmt.Sprintf("%s%s%s%s%s", space, string(alphabet[x]),internos,string(alphabet[x]), space) ) 
            }else{
                
                	finalText = append(finalText ,fmt.Sprintf("%s%s%s", space, string(alphabet[x]), space) )
            }
         
        }

       
    

    return strings.Join(finalText, "\n"), nil
}
