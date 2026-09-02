package queenattack
import "errors"

var firstPosition = map[string]int{
        "a":0,
        "b":1,
        "c":2,
        "d":3,
        "e":4,
        "f":5,
        "g":6,
        "h":7,
        
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
    
    //white validation 
	if len(whitePosition) != 2  || whitePosition[0] <'a' || whitePosition[0] > 'h' || whitePosition[1]-'0'<= 0 || whitePosition[1]-'0' >8{
        return false , errors.New("Error..")
    }
    //black validation
    if len(blackPosition) != 2 || blackPosition[0] <'a' || blackPosition[0] > 'h' || blackPosition[1]-'0'<= 0 || blackPosition[1]-'0' >8{
        return false , errors.New("Error..")
    }
    
    whitePositioncolum1:= firstPosition[string(whitePosition[0])]
    whitePositionrow1:= 8 - int(whitePosition[1]-'0')
	
    whitePositioncolum2:= firstPosition[string(blackPosition[0])]
    whitePositionrow2:= 8 - int(blackPosition[1]-'0')

    if whitePositionrow1 == whitePositionrow2 && whitePositioncolum1==whitePositioncolum2{
        return false , errors.New("Error..")
    }
    

    if (whitePositionrow1 == whitePositionrow2) || (whitePositioncolum1 == whitePositioncolum2) {
        return true , nil
    }

   
    if diferenciaAbsoluta(whitePositionrow1 , whitePositionrow2) == (diferenciaAbsoluta(whitePositioncolum1, whitePositioncolum2)){
        return true , nil
    }
    
    return false , nil
}

func diferenciaAbsoluta(a,b int)(int ){
    if a>b{
    	return a-b
    }
    return b-a
        
    
}
