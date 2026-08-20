package primefactors

func Factors(n int64) []int64 {
    arr:= []int64{}
	i := 2
	for n > 1 {
        if n % int64(i) == 0 {
            arr = append(arr, int64(i))
            n = n / int64(i)
        } else {
            i++
        }
}
   
    return arr
}
