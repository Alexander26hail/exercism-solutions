package birdwatcher
import ("fmt")
// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    total:=0
	for _ , birth := range  birdsPerDay {
        total+=birth
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total:=0
    totaldias:=(7*week)-7
    totalbirths:= []int {}
    fmt.Printf("valor... %d" ,totaldias )
    for i:=totaldias; i<totaldias+7; i++{
        fmt.Printf("valor... %v" , birdsPerDay[i] )
        totalbirths= append(totalbirths,birdsPerDay[i])
    }
    total= TotalBirdCount(totalbirths)
	return total
    
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	
    
    for i:=0; i<len(birdsPerDay); i ++{
        
        if i%2==0 {
            birdsPerDay[i]++
         	
        }
    }
    return birdsPerDay
}
