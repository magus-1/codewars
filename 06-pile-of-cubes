package kata

import (
	"fmt"
	"math"
)

func FindNb(m int) int {
	// input:
	//  m = total volume
	// output:
	//  n = total cubes
	// Cubes are stacked. Top cube's volume is 1, each lower cube is (n-1)^3, with n being current level
	// Find sumN for the given m. Return -1 if it doesn't match a possible combination of sumN
	var currentM float64 = 0
	target := float64(m)
	var n float64 = 0

	for currentM < target {
		n += 1
		currentM = currentM + math.Pow((n-1), 3)
		switch {
		case currentM < float64(target):
			continue
		case currentM == target:
			fmt.Printf("CurrentM: %v, target: %v, n: %v", currentM, target, n)
			return int(n - 1)
		}
	}
	fmt.Printf("CurrentM: %v, target: %v, n: %v", currentM, target, n)
	return -1
}
