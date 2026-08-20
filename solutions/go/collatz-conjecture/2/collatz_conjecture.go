package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
	
    var counter = 0
    if n <= 0{
        return 0, errors.New("Error...")
    }
    for n != 1 {
        if n % 2 == 0{
            n = n /2
            counter ++
        }else{
            n = n * 3 +1
            counter ++ 
        }
        
    }

    return counter, nil 
    
    
}

