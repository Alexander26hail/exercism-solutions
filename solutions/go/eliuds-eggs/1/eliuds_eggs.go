package eliudseggs

func EggCount(displayValue int) int {
    count:=0
	for displayValue> 0{
        
        if displayValue % 2 == 1{
            count ++
        }
    	displayValue = displayValue /2
    }
    return count
    
	

}
