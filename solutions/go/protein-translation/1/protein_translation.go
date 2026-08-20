package proteintranslation
import "errors"
var protein = map[string]string{
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP",
    "UAG": "STOP",
    "UGA": "STOP",
}
var     ErrStop = errors.New("ErrStop")
var 	ErrInvalidBase = errors.New("ErrInvalidBase")

func FromRNA(rna string) ([]string, error) {
	codon:= []string{}
    
	for i:=0; i<len(rna); i+=3{
        if i+3 > len(rna) { 
            return codon, ErrInvalidBase
        }
        value:= rna[i:i+3]
       
        response,err := FromCodon(value)
        
        if err == ErrStop {
    		return codon, nil  // ← STOP → devuelve lo acumulado sin error
        }
        if err != nil {
            return codon, err  // ← otro error → propaga
        }
        codon= append(codon,response)
        
    }
    return codon, nil
}

func FromCodon(codon string) (string, error) {
	if len(codon) !=3{
        return "", ErrInvalidBase
    }
    
    condonval, validation := protein[codon]
    if !validation {
        
        return "", ErrInvalidBase
    }
    if condonval == "STOP"{
        return condonval, ErrStop
    }else{
        return condonval, nil
    }
    
}
