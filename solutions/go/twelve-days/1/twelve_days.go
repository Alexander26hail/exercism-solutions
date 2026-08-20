package twelvedays
import "fmt"
var numberToLetter = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth", 
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(i int) string {
    verse:=""
    letter:=""
	if i >=1 && i <=12{
        
            switch {
                case i==1:
                	letter += "a Partridge"
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: a Partridge in a Pear Tree." , string(numberToLetter[i]) )
                case i==2:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==3:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==4:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==5:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==6:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==7:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
				case i==8:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==9:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==10:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==11:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))
                case i==12:
                	letter += "two Turtle Doves, "
                	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me: twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree." , string(numberToLetter[i]))

                
            }
            
        
    }
    return verse
}

func Song() string {
    versefinal:=""
	for i := 1; i <= 12; i++ {
        if i ==12{
            versefinal+= Verse(i) 
        }else{
            versefinal+= Verse(i) + "\n"
        }
        
        
    }
    return versefinal
}
