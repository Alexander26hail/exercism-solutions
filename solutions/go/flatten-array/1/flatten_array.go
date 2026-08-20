package flattenarray

func Flatten(nested any) []any {
   arrayFlatten:= []any{}
	if arrayFlatten == nil{
        return arrayFlatten
    }
    elem , isarray := nested.([]any)
    if isarray {
        for _ , elemArray := range elem{
            if elemArray!= nil{
                flat := Flatten(elemArray)
                arrayFlatten = append(arrayFlatten , flat...)
            }
        }
    }else{
        arrayFlatten = append(arrayFlatten , nested)
    }

    return arrayFlatten
}
