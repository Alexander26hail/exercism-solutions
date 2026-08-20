package largestseriesproduct
import "fmt"
import "errors"
//import "strings"
import "unicode"
func LargestSeriesProduct(digits string, span int) (int64, error) {

    var inicial int64 = 0
    var array = []string{}
	if len(digits) < span || span < 0    {

        return inicial, errors.New("Error...")
    }
    
    for i := 0; i <= len(digits) - span; i++{
		if unicode.IsLetter(rune(digits[i])){
            return inicial, errors.New("Error...")
        }
        array = append(array, digits[i:i+span])
        fmt.Printf("value : %s", digits[i:i+span])
       
        
        
    }
    inicial= 0
    for _ , arr := range array {
        
        var multip int64 = 1
        for _, arr2 := range arr{
            multip*=  int64(arr2 - '0')
             
        }
        if inicial< multip{
            fmt.Printf("historico: %d   multi: %d ", inicial ,multip)
            inicial = multip
    	 }
       
    }
    

    return inicial , nil 
}
