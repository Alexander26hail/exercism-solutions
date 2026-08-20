package rnatranscription
import "strings"
var RNA = map[string]string{
    "G":"C",
    "C":"G",
    "T":"A",
    "A":"U",
}
func ToRNA(dna string) string {
	dna = strings.ToUpper(dna)
	value:=""
    for _ , i := range dna { 
        rna , exist := RNA[string(i)]
    	if exist{
            value+=rna
        }
    } 
    return value
}
