package dndcharacter
import "math/rand"
import "math"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
 	arrayDice := []int{}
	arrayDice = append(arrayDice,rand.Intn(6)+1) 
	arrayDice = append(arrayDice,rand.Intn(6)+1) 
    arrayDice = append(arrayDice,rand.Intn(6)+1) 
    arrayDice = append(arrayDice,rand.Intn(6)+1) 
    skills:=0
    min:=arrayDice[0]
    for _ , i := range arrayDice {
        if min > i {
            min = i
        }
        skills+= i
        
    }
    return skills - min

    
	

    
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    Consti:= Ability()
	return Character{
        Strength     :Ability(),
    	Dexterity    :Ability(),
    	Constitution :Consti,
    	Intelligence :Ability(),
    	Wisdom       :Ability(),
    	Charisma     :Ability(),
    	Hitpoints    :10 + Modifier(Consti),
    }
}
