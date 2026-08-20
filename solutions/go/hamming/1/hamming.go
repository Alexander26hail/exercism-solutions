package hamming
import "errors"
func Distance(a, b string) (int, error) {
    sum :=0 
	if len(a) != len (b){
        return sum , errors.New("Error mayor al valor")
	}

    for i , c := range a {
        if c != rune(b[i]) {
            sum++
        }
    }
    return sum,nil
}
