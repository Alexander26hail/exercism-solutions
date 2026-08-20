package etl
import("fmt")
import("strings")
func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)

    for puntos, letras := range in {
        fmt.Printf("imprimimos %d", puntos)

        for _, letra := range letras {
            fmt.Printf("imprimimos %s ", letra)
            lowerCase:= strings.ToLower(letra)
            result[lowerCase]= puntos
        }
    }
    
    return result
}
