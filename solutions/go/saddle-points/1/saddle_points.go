package saddlepoints
import "strings"
import "strconv"
import "errors"
// Define the Matrix and Pair types here.
type Matrix struct{
    rows [][]int
}
type Pair struct{
    row int
    column int
}
func New(s string) (*Matrix, error) {
   
	SplitText:= strings.Split(s, "\n")
    rows:=[][]int{}
    for _ , spaceRow := range SplitText{
        
        numeroTexto := strings.Fields(spaceRow)
        
        actualRow:= []int{}
        
        for _ , elem := range numeroTexto{
            num, err := strconv.Atoi(elem)
        	if err != nil {
        		return nil, err
            		
        	}
			actualRow = append(actualRow , num)
            
        }
        if len(rows)!=0 && len(rows[0])<len(actualRow){
            return nil, errors.New("error") 
        }
        
        rows = append(rows , actualRow)
    }
   

    return &Matrix{rows: rows }, nil
            	
}

func (m *Matrix) Saddle() []Pair {
	result:= []Pair{}
    colum:= m.Cols()

    for r , _ := range m.rows{
        for c  , _ := range colum {
            value:= m.rows[r][c]
            var isUpperRow = true
			for _, otro1 := range m.rows[r] {
                if otro1 > value {
                    isUpperRow = false
                    break
                }
            }

            isUpperColum:= true
            for _, otro2 := range colum[c] {
                if otro2 < value {
                    isUpperColum = false
                    break
                }
            }

            if isUpperRow  && isUpperColum {
                result = append(result, Pair{r + 1, c + 1})
            }
        }
        
    }
    return result
}

func (m Matrix) Cols() [][]int {
    ColumNumeros:=[][]int{}

	for number , _ := range  m.rows[0] {
        listaActual:= []int{}
        for number2 , _ := range m.rows {
            listaActual = append(listaActual , m.rows[number2][number] )
        }
        ColumNumeros = append(ColumNumeros , listaActual) 
    }
    return ColumNumeros
}