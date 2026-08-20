package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/02/2006 15:04:05"
	clock, _ := time.Parse(layout, date)
    
	return clock
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	 layout := "January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)
    
    // Si la fecha de la cita (t) es anterior a "ahora mismo", ya pasó
    return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    
    t, _ := time.Parse(layout, date)
    

    return t.Hour() >= 12 && t.Hour() <18
         
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    format:= "Monday, January 2, 2006, at 15:04."
    t, _ := time.Parse(layout, date)

    return "You have an appointment on " + t.Format(format)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    
	
    t := time.Now().Year()
	anniversary := time.Date(t, time.September, 15, 0, 0, 0, 0, time.UTC) 
    return anniversary
    
}
