package lasagna

// PreparationTime calculates how long it will take to prepare the lasagna
// given an average preparation time per layer and a number of layers.
func PreparationTime(layers []string, averageLayerPrepTime int) int {
	if averageLayerPrepTime == 0 {
		averageLayerPrepTime = 2
	}

	return averageLayerPrepTime * len(layers)
}

// Quantities calculates the total amount of noodles and sauce for the
// noodle and sauce layers of a lasagna.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return
}

// AddSecretIngredient adds the unknown secret ingredient from friendsList to myList.
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales a recipe written for 2 portions to any given number of portions
func ScaleRecipe(measurements []float64, portions int) []float64 {
	scaled := make([]float64, len(measurements))

	for i, measurement := range measurements {
		scaled[i] = measurement * .5 * float64(portions)
	}

	return scaled
}
