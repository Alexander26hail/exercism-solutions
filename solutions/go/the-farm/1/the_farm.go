package thefarm
import "errors"
import "fmt"
var Messagenegative = errors.New("there are no negative cows")
var Messagefood = errors.New("no cows don't need food")


// TODO: define the 'DivideFood' function
func  DivideFood(fod FodderCalculator,  cows int) (float64 , error){
    FoddedAmo ,  err :=fod.FodderAmount(cows)
    if err !=nil{
        return 0.0 , err 
    }
    FatterFact, err2:= fod.FatteningFactor()
    if err2 != nil{
        return 0.0 , err2
    }
    return (FoddedAmo*FatterFact)/float64(cows), nil
}
// TODO: define the 'ValidateInputAndDivideFood' function

func  ValidateInputAndDivideFood(fod FodderCalculator,  cows int) (float64 , error){
    if cows>0{
        DivFood, err := DivideFood(fod , cows)
        if err != nil{
            return 0.0 , err
        }
        return DivFood , nil
    }
	
    return 0 , errors.New("invalid number of cows")

    

}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error{
    if cows <0{
        return fmt.Errorf("%d cows are invalid: %w", cows , Messagenegative)
    }
    if cows == 0{
        return fmt.Errorf("%d cows are invalid: %w", cows , Messagefood)
    }
    return  nil
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
