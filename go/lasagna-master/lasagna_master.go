package lasagna

func PreparationTime(layers []string, average int) int {
	// provide a default value of 2 if 0 is passed
	if average == 0 {
		average = 2
	}
	return len(layers) * average
}

func Quantities(layers []string) (int, float64) {
	noodlesQuantity := 0
	sauceQuantity := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesQuantity += 50
		case "sauce":
			sauceQuantity += 0.2
		}
	}
	return noodlesQuantity, sauceQuantity
}

func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64
	for _, quantity := range quantities {
		scaledQuantities = append(scaledQuantities, quantity/2*float64(portions))
	}
	return scaledQuantities
}
