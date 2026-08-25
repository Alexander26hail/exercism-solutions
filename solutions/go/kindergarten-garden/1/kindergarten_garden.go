package kindergartengarden
import "strings"
import "slices"
import "errors"
// Define the Garden type here.
type Garden struct{
    fila1 string 
    fila2 string
    childen []string
}
// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`
//
// If the children argument is empty, use the list of children defined in the instructions.
var ChildensList = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"}

var planted = map[string]string{
    "G":"grass",
    "C":"clover",
    "R":"radishes",
    "V":"violets",
}
// If it is not empty, use the given value.

func NewGarden(diagram string, children []string) (*Garden, error) {
	partes := strings.Split(diagram , "\n")
    if len(partes) < 3 || len(partes[1]) != len(partes[2]) ||  len(partes[1])%2 == 1 {
    	return nil, errors.New("Error")
	}
    fila1:=partes[1]
    fila2:=partes[2]

    for _ , i := range fila1{
        _ , exist1 := planted[string(i)]
        if !exist1 {
            return nil, errors.New("Error")
        }
    }
    
	var listaFinal = []string{}
    if len(children) ==0{
        listaFinal= ChildensList
       
    }else{
        copyChildren := slices.Clone(children)
        slices.Sort(copyChildren)
        listaFinal=copyChildren
        vistos := make(map[string]bool)
		for _, v := range listaFinal {
            if vistos[v] {
                return nil, errors.New("Error")
        	}
            vistos[v] = true
    	}
        
    }

    return &Garden{
        fila1:fila1 , 
    	fila2:fila2 ,
    	childen :listaFinal ,
    } ,  nil
	
    
    
}

func (g *Garden) Plants(child string) ([]string, bool) {
	position := -1
    for i , name := range g.childen{
        if name == child{
            position = i 
            break
        }
    }
    if position == -1 {
   		 return nil, false
	}
    position = position * 2 
    var plants = []string{}
    plants = append(plants,  planted[string(g.fila1[position])])
    plants = append(plants , planted[string(g.fila1[position+1])])
    plants = append(plants , planted[string(g.fila2[position])])
    plants = append(plants , planted[string(g.fila2[position+1])])

    return plants , true
     
}
