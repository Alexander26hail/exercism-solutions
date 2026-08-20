package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial

    for _, list:= range s {
        result = fn(result,list )
    }
    return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial

    for i:= len(s) -1; i>=0; i--{
        result= fn(s[i],result)
    }
    return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	IntNew:=  []int {}
	for _, lst := range  s {
        if fn(lst){
            IntNew = append(IntNew,lst )
        }
    }
    return IntNew
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
 IntNew := []int{}
    for _, lst := range s {
        IntNew = append(IntNew, fn(lst))
    }
    return IntNew
}

func (s IntList) Reverse() IntList {
	IntNew := []int{}
    for i:=len(s) -1; i>=0; i--{
        IntNew = append(IntNew, s[i])
    }
    return IntNew
}

func (s IntList) Append(lst IntList) IntList {
	

    for _, lst2 := range lst {
        s = append(s,lst2 )
    }
    return s 
}

func (s IntList) Concat(lists []IntList) IntList {
	 for _, lst2 := range lists {
        s = append(s,lst2...)
    }
    return s 
}
