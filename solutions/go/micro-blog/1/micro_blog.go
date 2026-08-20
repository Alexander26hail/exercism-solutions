package microblog

func Truncate(phrase string) string {
	count:= 0 
    words:=""
	for _ , phr := range phrase{
        
        if count != 5 {
            words +=  string(phr)
            count++
        }
    }
    return words
}
