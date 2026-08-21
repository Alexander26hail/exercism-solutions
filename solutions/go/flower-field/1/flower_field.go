package flowerfield
import "fmt"
// Annotate returns an annotated board
func Annotate(board []string) []string {
    filas:= len(board)
	if filas == 0{
        return board
    }
    colum:= len(board[0])
    result:= []string{}
    for r:=0; r<=filas-1; r++{
        nuevaFila:=""
        for c:=0; c<=colum-1; c++{
            if board[r][c] == '*'{
                nuevaFila += "*"
            }else{
                contador := 0
                for dr:=-1; dr<=1; dr++{
                    for dc:=-1; dc<=1; dc++{
                        if dr==0 && dc==0{
                            continue
                        }
                        vf:= dr + r
                        vc:= dc + c

                        if (vf >= 0) && (vf < filas) &&  (vc >= 0) && (vc < colum) {
                            if board[vf][vc] ==  '*'{
                                contador++
                            }
                        }
                        
                        
                    }
                }
                if contador == 0{
                    nuevaFila += " "
                }else{
                    nuevaFila += fmt.Sprintf("%d", contador)  
                }
                
            }
            
        }
        result = append(result , nuevaFila )
    }
    
    
    return result
}
