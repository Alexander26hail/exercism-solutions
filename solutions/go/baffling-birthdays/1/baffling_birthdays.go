package bafflingbirthdays

import "time"
import "math/rand/v2"
//import "fmt"
func SharedBirthday(dates []time.Time) bool {
   
	for i , date := range dates{
        
        for x:=i+1; x<len(dates); x++{
            
            if date.Day() == dates[x].Day()  && date.Month() == dates[x].Month(){
            	return true 
        	}
        }
        
            
        
    }
    return false
}

func RandomBirthdates(size int) []time.Time {
	randomDates:= make([]time.Time, size)

    for i := 0; i < size; i++ {
        day := rand.IntN(31) + 1            // 1-28
        month := time.Month(rand.IntN(12) + 1)  // 1-12
        year := rand.IntN(50)*2 + 1971       // 1971-1973
        randomDates[i] = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}
    return randomDates
}

func EstimatedProbability(size int) float64 {
	simulations:=10000
    collitions:=0
	
    for i:=0; i<simulations; i++{
        if SharedBirthday(RandomBirthdates(size)) == true{
            collitions++
        }
    }
    return (float64(collitions)/ float64(simulations) ) * 100
}
