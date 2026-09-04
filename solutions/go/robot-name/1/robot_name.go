package robotname
import "math/rand" 
import "fmt"
import "errors"
// Define the Robot type here.
type Robot struct{
    Robotname string
}
var listName = map[string]bool{}
func (r *Robot) Name() (string, error) {
    if len(listName) >= 676000{
        return "" , errors.New("Error")
    }
    if r.Robotname !=""{
         
        return r.Robotname , nil
    }
    for {
        randNumber:= rand.Intn(1000-1)+1
       	letter:=""
    	
        for i:=0; i<2; i++{
            randLetter := rand.Intn(90-65+1) + 65
            letterRune := rune(randLetter)
            letter += string(letterRune)
            
        }
        name:= fmt.Sprintf("%s%03d", letter,randNumber)
        
        _ , exist := listName[name]
        if !exist{
            r.Robotname = name
            listName[r.Robotname]= true
            return r.Robotname , nil
        }
        
    	
    }
	
    return r.Robotname , nil
  
}

func (r *Robot) Reset() {
	r.Robotname = ""
}
