package squareroot
import "errors"
func SquareRoot(number int) (int, error) {
    if number > 0{
        for i:=1; i* i <=number; i++{
            if i*i == number {
                return i , nil
            }
    	}
    }
	return 0 , errors.New("error")
}
