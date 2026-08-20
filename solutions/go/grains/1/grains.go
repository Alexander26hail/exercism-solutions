package grains
import "errors"
func Square(number int) (uint64, error) {
    power:=1
	if number > 0 && number <=64 {
        for i:=1; i<number; i++{
            power*=  2 
        }
        return uint64(power) , nil 
    }
    return 0 , errors.New("Error")
    
    
}

func Total() uint64 {
	var total uint64
	for i:=1; i<=64; i++{
        totalbefore , _ :=Square(i)
        total+=totalbefore
    }
    return total
}
