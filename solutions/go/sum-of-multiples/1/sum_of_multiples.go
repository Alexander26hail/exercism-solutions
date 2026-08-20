package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	array := []int{}
	for _ , div := range  divisors{
        if div <=0{
            continue 
        }
        for multi:=div; multi<limit; multi+=div{
            array= append(array, multi)
        }
    }

    arrayfinal:= RemoveDuplicates(array)
	valortotal:= 0 
    for _, sum := range arrayfinal{
        valortotal += sum
    }
    return  valortotal
}
func RemoveDuplicates(arr []int ) []int {
    visto := make(map[int] bool)
    arrayFinal:= []int{}

    for _, arrs:= range arr{
        if !visto[arrs]{
            arrayFinal = append(arrayFinal, arrs)
            visto[arrs]= true 
        }
    }

    return arrayFinal
}