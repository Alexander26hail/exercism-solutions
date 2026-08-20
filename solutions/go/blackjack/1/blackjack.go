package blackjack
var cards = map[string]int {
    "ace":11,
    "two":2,
    "three":3,
    "four":4,
    "five":5,
    "six":6,
    "seven":7,
    "eight":8,
    "nine":9,
    "ten":10,
    "jack":10,
    "queen":10,
    "king":10,
    "other":0,
    
}
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	
    value, _ := cards[card]

    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	valuecar1  := ParseCard(card1)
    valuecar2  := ParseCard(card2)
    sum := cards[card1] + valuecar2 
    dealer:= ParseCard(dealerCard)

	switch {
        case valuecar1 == 11 &&  valuecar2== 11:
        	return "P"
        case sum == 21 :
        	if dealer>= 10  && dealer<= 11 {
                return "S"
            }
        	return "W"
        	
        
        case sum >= 17 && sum <= 20:
			return "S"
        
        case sum >= 12 && sum <= 16 :
        	if dealer >= 7{
                return "H"
            }
        	return "S"
        	
        
        case sum <=11:
        	return  "H"
        default: 
        	return ""
    }    
	
}
