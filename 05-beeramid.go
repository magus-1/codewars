package kata

import (
	"math"
)

func Beeramid(bonus int, price float64) int {
	// your code here
	var pool = float64(bonus) // remaining money
	for level := 0; pool >= 0; level++ {
		cost := math.Pow(float64((level+1)), 2) * price
		if cost > pool {
			return level
		}
		pool -= cost
	}
	return 0
}
