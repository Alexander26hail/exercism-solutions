package customset
import "strings"
import "slices"
// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.

// Define the Set type here.
type Set map[string]bool
func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
    var SetNew= New()
	for _ , elem := range l {
        SetNew.Add(elem)
    }
    return SetNew
}

func (s Set) String() string {

    array:= []string{}
	for elem  := range s { 
    	array = append(array , "\""+elem + "\"")
    }
    slices.Sort(array)
    
    return "{" + strings.Join(array , ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_ , exist := s[elem]
    return exist
}

func (s Set) Add(elem string) {
    
	s[elem]= true
}

func Subset(s1, s2 Set) bool { 
  
	for elem := range s1 {
        if !s2.Has(elem) {
            return false

        }
    }
    return true
}

func Disjoint(s1, s2 Set) bool {
	for elem := range s1 {
        if s2.Has(elem){
            return false
        }
    }
    return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2){
        return false
    }
    for elem  := range s1{
        if !s2.Has(elem){
            return false
        }
    }
    return true 
}

func Intersection(s1, s2 Set) Set {
    s3 := Set{}
	
    for elem := range s1{
        if s2.Has(elem) {
            s3.Add(elem)
        }
        
    }
    return s3
    
    
}

func Difference(s1, s2 Set) Set {
    s3:= Set{}
	
    for elem := range s1{
        if !s2.Has(elem){
            s3.Add(elem)
        }
    }
    return s3
}

func Union(s1, s2 Set) Set {
    s3:= Set{}

    
    for elem:= range s1{
        s3.Add(elem)
    }
    for elems2:= range s2{
        s3.Add(elems2)
    }
    return s3
}
