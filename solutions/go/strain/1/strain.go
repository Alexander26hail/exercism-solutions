package strain
type Predicado[T any] func(T) (bool)
// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](lista []T, p  Predicado[T] ) []T{
    ListResult:= []T{}
    for _ , element:= range lista{
        if p(element){
            ListResult = append(ListResult , element)
        }
    }
    return ListResult
}
func Discard[T any](lista []T, p  Predicado[T]) []T{
    ListResult:= []T{}
    for _ , element:= range lista{
        if !p(element){
            ListResult = append(ListResult , element)
        }
    }
    return ListResult

}
