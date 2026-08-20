package binarysearch
import "slices"

func SearchInts(list []int, key int) int {
	slices.Sort(list[:])
    
    left := 0
	right := len(list) - 1

    for left <= right {
        mid := (left + right) / 2
       
        if key == list[mid] {
            return mid  
        }
        if key > list[mid] {
        	left = mid + 1   
        } else {
            right = mid - 1  
        }
        
    }
    return -1
    
}

