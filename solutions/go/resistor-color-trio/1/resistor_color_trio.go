package resistorcolortrio
import "fmt"
// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

var mapColors = map[string]int{
    "black": 0,
    "brown": 1,
    "red": 2,
    "orange": 3,
    "yellow": 4,
    "green": 5,
    "blue": 6,
    "violet": 7,
    "grey": 8,
    "white": 9,
}

func Label(colors []string) string {
	number := mapColors[colors[0]] *10  + mapColors[colors[1]]
    lastColor := mapColors[colors[2]]

    for i:=0; i<lastColor; i++{
        number = number *10
    }
	
	sufix := []string{"ohms", "kiloohms", "megaohms", "gigaohms"}
	index:=0
    for number >= 1000 && number%1000 == 0{
        number= number/1000
        index++
    }
	return fmt.Sprintf("%d %s", number ,sufix[index] )
  
    /*
    if number<1000{ 
        return fmt.Sprintf("%d ohms", number)
    }else{
        switch {
            case number/1000<1000:
            	return fmt.Sprintf("%d kiloohms", number/1000)
            case number/1000000<1000:
				return fmt.Sprintf("%d megaohms", number/1000000)
            case number/1000000000<1000:
            	return fmt.Sprintf("%d gigaohms", number/1000000000)
        }

      return "" 
    }*/

 
          
}
