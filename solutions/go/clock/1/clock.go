package clock
import "fmt"
// Define the Clock type here.
type Clock struct {
    hour int
    min int
    
}
func New(h, m int) Clock {
    total:= h*60+m
    total= ((total%1440) +1440 )%1440
    return Clock{hour:total/60,min:total%60}
    
	
}

func (c Clock) Add(m int) Clock {
  
    total:= c.hour * 60 + + c.min + m
   total= ((total%1440) +1440 )%1440
	return Clock{hour:total/60,min:total%60}
}

func (c Clock) Subtract(m int) Clock {
    total:= c.hour * 60 + + c.min - m
    total = ((total%1440) +1440 )%1440
	return Clock{hour:total/60,min:total%60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour , c.min)
}
