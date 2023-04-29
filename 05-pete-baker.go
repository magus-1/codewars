package kata

func Cakes(recipe, available map[string]int) int {
	maxCakes, maxIngredient := int(^uint(0)>>1), 0
	for k, v := range recipe {
		maxIngredient = available[k] / v
		if maxIngredient < maxCakes {
			maxCakes = maxIngredient
		}
	}
	return maxCakes
}
