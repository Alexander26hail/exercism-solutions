package matchingbrackets

func Bracket(input string) bool {
    pila:=[]rune{}
  	parejas := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
	}
    for _ , elem := range input{
       if elem == '(' || elem == '[' || elem == '{'{
           pila = append(pila,elem)
       }
        
		lastelement , exist := parejas[elem]
		
        if exist {
			if len(pila)==0{
                return false
            }
            cima := pila[len(pila)-1]
            
            if cima != lastelement{
                return false
            }
            pila = pila[:len(pila)-1]
            
        }
        

        
        
        
    }

    return len(pila)==0

}
