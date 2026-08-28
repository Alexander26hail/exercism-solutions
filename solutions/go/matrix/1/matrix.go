package saddlepoints
import "strings"
import "strconv"
import "slices"
import "errors"
// Define the Matrix type here.
type Matrix struct  {
    rows	[][]int
}
func New(s string) (Matrix, error) {
    filaNumeros:=[][]int{}
    normText:= strings.Split(s, "\n")
    for _ , elem:= range normText { 
        spacesText:= strings.TrimLeft(string(elem), " ")
        numeroTexto:= strings.Split(spacesText, " ")
        listaActual:= []int{}
		for _ , numText:=range numeroTexto{
            texto , err := strconv.Atoi(numText)
            if err!=nil{
                return Matrix{
            		rows: [][]int{},
        		}, err
            }
            listaActual =  append(listaActual, texto)
            
        }
        if len(filaNumeros)!=0 && len(filaNumeros[0])<len(listaActual){
            return Matrix{
            		rows: filaNumeros,
        		}, errors.New("error") 
        }
        filaNumeros = append(filaNumeros , listaActual )
    }
    
	return Matrix{
        rows: filaNumeros,
    }, nil 
}

// Cols and Rows must return the results without affecting the matrix.
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

func (m Matrix) Rows() [][]int {
    ColumNumeros:=[][]int{}
    for _ , elem := range m.rows  {
        ColumNumeros = append(ColumNumeros , slices.Clone(elem))
    }
	return ColumNumeros
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.rows){
        return false
    }
    if col < 0 || col >= len(m.rows[row]){
        return false
    }
    m.rows[row][col]= val

    return true
}
