package anagram
import "strings"
import "slices"
import("fmt")
func Detect(subject string, candidates []string) []string {
	slicerunner := []rune(strings.ToUpper(subject))
	slices.Sort(slicerunner) 

    newsubject:= string (slicerunner)
    candidatesNorm:= ""
    
    final := []string{}
    for i:=0; i<len(candidates); i++{
         
        slicecandidates:=[]rune(strings.ToUpper(candidates[i]))
        slices.Sort(slicecandidates)
        candidatesNorm= string (slicecandidates)
		fmt.Println("Debug message " + candidatesNorm + " otro " + subject )
        if ( strings.ToUpper(candidatesNorm) == strings.ToUpper(newsubject) &&  strings.ToUpper(candidates[i]) != strings.ToUpper(subject)){
            final = append(final,candidates[i])
        }
        

        
        
    }
    return final
}
