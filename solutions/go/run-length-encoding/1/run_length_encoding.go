package runlengthencoding
import "strings"
import "strconv"
func RunLengthEncode(input string) string {
    if input==""{
        return ""
    }
    anterior:= input[0]
    
    var result strings.Builder
    result.Grow(len(input))
    
    count := 1 
    
	for i:=1; i<len(input); i++{
		if anterior == input[i]{
            count++
        }else{
            volcar(&result , count, anterior)
            anterior = input[i]
            count = 1
        }
    }
    
   	volcar(&result , count, anterior)
    
    return result.String() 
}

func volcar(builder *strings.Builder, count int, caracter byte) {
	if count == 1 {
		builder.WriteByte(caracter) 
	} else {
		builder.WriteString(strconv.Itoa(count))
		builder.WriteByte(caracter) 
	}
}

func RunLengthDecode(input string) string {
    var result strings.Builder
    number := 0
	if input == ""{
        return input
    }
    result.Grow(len(input) * 2)
    
    for _ , elem := range input{
        if elem>= '0' && elem <= '9'{
            number = number*10 + int(elem-'0')
        }else{
            if number != 0{
                for i:=0; i<number; i++{
                    result.WriteRune(elem)
                }
                number = 0
            }else{
                result.WriteRune(elem)
            }
        }
    }
    

    return result.String()
}
