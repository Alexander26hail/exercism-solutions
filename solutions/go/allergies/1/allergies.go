package allergies
var allergiesMAP = []struct {
    bit  uint
    name string
}{
    {128, "cats"},
    {64, "pollen"},
    {32, "chocolate"},
    {16, "tomatoes"},
    {8, "strawberries"},
    {4, "shellfish"},
    {2, "peanuts"},
    {1, "eggs"},
}
func Allergies(allergies uint) []string {
    var array []string
	
    for _ , text := range allergiesMAP{
        if allergies & text.bit != 0{
            array = append(array, text.name )
            
        }
    }
    return array
}

func AllergicTo(allergies uint, allergen string) bool {

    for _ , aler := range allergiesMAP {
        if allergies & aler.bit !=0 && aler.name == allergen {
            return true
        }
    }
    return false 
}
