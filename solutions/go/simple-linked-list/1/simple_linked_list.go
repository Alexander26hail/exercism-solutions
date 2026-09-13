package simplelinkedlist
import "errors"
type Element struct {
	Value int           
	Next  *Element
}

type List struct {
	Head *Element      
}

func New(elements []int) *List {
    NewList := List{}
	for _ , elem := range elements {
        NewList.Push(elem)
        
    }
    return &NewList
}

func (l *List) Size() int {
	count := 0 
    actual:= l.Head
    for actual != nil {
        count ++
        actual = actual.Next
    }
    return count
}

func (l *List) Push(element int) {
	ElementNuevo := &Element{
        Value:element,
    }
    ElementNuevo.Next = l.Head
    l.Head = ElementNuevo
}

func (l *List) Pop() (int, error) {
	if l.Head == nil{
        return 0 , errors.New("list is empty")
    }
    header := l.Head.Value
    l.Head = l.Head.Next
    return header , nil
    
}

func (l *List) Peek() (int, error) {
	if l.Head == nil  {
        return 0 , errors.New("list is empty")
    }
    return l.Head.Value , nil 
}

func (l *List) Array() []int {
    array:= []int{}
    arrayReverse := []int{}
    actual:= l.Head
    for actual != nil{
        array = append(array, actual.Value)
        actual = actual.Next
    }
	for i:=len(array)-1; i>=0; i--{
        arrayReverse = append(arrayReverse , array[i])
    }
    
	return arrayReverse
}

func (l *List) Reverse() *List {
	newReverse := List{}
    actual := l.Head
    for actual != nil{
        newReverse.Push(actual.Value)
        actual = actual.Next
    }
    return &newReverse
}
