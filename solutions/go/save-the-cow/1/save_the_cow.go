package savethecow
import "errors"
import "strings"

type Game struct{
    name string
    fail int
    listname map[rune]bool
}


func NewGame(word string) *Game {
	return &Game{	name: word,
                	fail:0,
                 	listname: map[rune]bool{},
                }
}

func (g *Game) Guess(r rune) error {
    if g.State() != "Ongoing"{
        if g.State() == "Win"   {
         	return errors.New("cannot guess after the game is won")
         }
        if g.State() == "Lose"   {
         	return errors.New("cannot guess after the game is lost")
         }
		return nil
        
    }
    _ , exist := g.listname[r]
    if !exist{
        g.listname[r]= true
    }else{
        g.fail++
        return nil
    }
    if !strings.ContainsRune(g.name , r ){
        g.fail++
    }
    return nil
}

func (g *Game) MaskedWord() string {
	result := ""
    for _ , elem:= range g.name{
        if g.listname[elem] {
            result += string(elem)
        }else{
            result += "_"
        }
        
    }
    
  return  result
}

func (g *Game) RemainingGuesses() int {
    result := 9-g.fail
    if result<=0{
        return 0
    }
    return result
}

func (g *Game) State() string {
    if g.fail >= 10{
        if  !strings.Contains(g.MaskedWord() , "_")  {
            return "Win"
        }
        return "Lose"
    }
    if strings.Contains(g.MaskedWord() , "_")  {
        return "Ongoing"
    }else{
        return "Win"
    }
	
}
