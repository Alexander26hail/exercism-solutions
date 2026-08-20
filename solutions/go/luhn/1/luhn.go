package luhn
import "fmt"
import "strings"


func Valid(id string) bool {
	
    clear:= strings.ReplaceAll(id, " ", "")
   
    fmt.Println("Hellow", clear)
	if len(clear)<=1{
        return false
    }
    valorTotal:= 0 
    contador:=0 
    for i:=len(clear)-1; i >= 0; i--{
        if clear[i] < '0' || clear[i] > '9' {
        	return false
    	}
        digito:= int(clear[i] - '0')
        if (contador % 2 == 1 ){
            digito = digito * 2 
            if digito > 9 {
                digito = digito -9 
            }
        }
        
        valorTotal +=  digito
		fmt.Println("Valor... ", digito)

        contador++
        
        
    }
	fmt.Println("Valor total ", valorTotal )
    if valorTotal % 10 == 0 {
        fmt.Println("paso aqui ")
        return true
    }
    fmt.Println("paso aqui es 1 ")

    return false 
    
}
