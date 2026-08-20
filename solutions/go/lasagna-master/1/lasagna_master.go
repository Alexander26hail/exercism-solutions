package lasagnamaster

type Layers []string

// PreparationTime calcula el tiempo total de preparación
func PreparationTime(layers Layers, avg int) int {
	if avg == 0 {
		avg = 2
	}
	return len(layers) * avg
}

// Quantities calcula la cantidad de fideos (int) y salsa (float64)
func Quantities(layers []string) (int, float64) {
	nuddle:= 0
    sauce:= 0.0
    for _, layers:= range layers{
        if layers== "noodles"{
            nuddle+= 50
        }
        if layers == "sauce"{
            sauce+=0.2
        }
    }
    return nuddle, sauce
    
}

// AddSecretIngredient reemplaza el último ingrediente de tu lista con el de tu amigo
func AddSecretIngredient(friendsList Layers, myList Layers) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe calcula las cantidades necesarias para el nuevo número de porciones
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var scaled []float64
	factor := float64(portions) / 2.0

	for _, amount := range amounts {
		scaled = append(scaled, amount * factor)
	}
	return scaled
}
