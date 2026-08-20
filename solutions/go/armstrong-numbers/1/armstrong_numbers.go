package armstrongnumbers
import "fmt"
import "math"
func IsNumber(n int) bool {
	conv:= fmt.Sprintf("%d", n )
    numDigits := len(conv)
    sum:=0
    for _, explit := range conv {
        digitValue := int(explit - '0')
        sum+= int(math.Pow(float64(digitValue), float64(numDigits)))
        
    }
    if sum == n {
        return true
    }
    return false
}
