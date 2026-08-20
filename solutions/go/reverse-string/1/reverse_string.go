package reversestring

func Reverse(input string) string {
	textFinal:= ""
    runes := []rune(input)
    for i:=len(runes) -1 ; i>=0; i -- {
        textFinal+= string(runes[i])
    }
    return textFinal
}
