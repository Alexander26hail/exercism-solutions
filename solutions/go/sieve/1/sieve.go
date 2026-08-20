package sieve

func Sieve(limit int) []int {
    var arrayInt []int
    isPrime := make([]bool, limit+1)
	if limit >=2{
        for i := 2; i <= limit; i++ {
            isPrime[i] = true
        }
        
        for i := 2; i <= limit; i++ {
            if isPrime[i] {
                arrayInt = append(arrayInt, i)
                for multiple := i * 2; multiple <= limit; multiple += i {
                    isPrime[multiple] = false
                }
            }
        }
    }
  
    return arrayInt
}
