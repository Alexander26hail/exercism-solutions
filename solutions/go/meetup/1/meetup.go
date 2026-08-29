package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
    First WeekSchedule = iota
    Second
    Third
    Fourth
    Last
    Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	gettingDays  := []int{}
    lastDay:= time.Date(year ,month+1 , 0 , 0 ,0 , 0 , 0, time.UTC)
	firstDay:=1
	if wSched == Teenth{
        firstDay=13
		lastDay = time.Date(year ,month , 19 , 0 ,0 , 0 , 0, time.UTC)
    }
    for i:=firstDay; i<=lastDay.Day(); i++{
        fecha := time.Date(year ,month , i , 0 ,0 , 0 , 0, time.UTC)
        if fecha.Weekday() == wDay{
            gettingDays = append(gettingDays , fecha.Day())
        }    
    }
    switch wSched{
        	case First:
				return gettingDays[0]
        	case Second:
				return gettingDays[1]
        	case Third:
				return gettingDays[2]
        	case Fourth:
				return gettingDays[3]
        	case Last:
				return gettingDays[len(gettingDays)-1]
			case Teenth:
        		return gettingDays[0]
    }
	return 0
}
