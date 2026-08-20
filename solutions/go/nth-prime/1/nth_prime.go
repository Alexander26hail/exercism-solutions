package nthprime
import "errors"
//import "fmt"
// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
    candidate:=2
    count:=0 
	if n<=0{
        return 0, errors.New("Error blank")
    }
    for count < n {
        if isprime(candidate){
            count ++ 
    	}

        candidate ++
        
    }
    
    return candidate -1  , nil
}
func isprime(n int)bool{
    for i:=2; i<n; i ++{
        if n % i == 0{
            return false
        }
        
    }
    return true 
}
