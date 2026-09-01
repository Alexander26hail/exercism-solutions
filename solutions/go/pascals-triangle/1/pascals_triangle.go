package pascalstriangle

func Triangle(n int) [][]int {
    result:= [][]int{}
	if n<=0{
        return result
    }
    for i:=0; i<=n-1; i++{
		arrayNumber:=[]int{}
        if i == 0{
            arrayNumber = append(arrayNumber , 1)
            result = append(result , arrayNumber)
            continue
        }
        for j:=0; j<=i; j++{
         	
            left:=0
            if j>0{
            	left = result[i-1][j-1]
        	}
            right:=0

            if j < len(result[i-1]){
                right = result[i-1][j]
            }
            arrayNumber = append(arrayNumber , right + left)
        }
		result = append(result, arrayNumber)
        
        
        
    }
    return result
}
