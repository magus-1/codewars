package kata

import (
	"fmt"
	"math"
)

func SumOfSquares(n uint64) uint64 {
	// Lagrange's four-square theorem states that at most the result will be 4
	// There are many algorithms that find the answer, but here's a simpler alternative using this theorem
	root := uint64(math.Sqrt(float64(n)))

	fmt.Printf("n: %v, root: %v, n&7: %v \n", n, root, n&7)
	// Step 1: check if n=1
	if root*root == n {
		return 1
	}

	// Step 2: check if n=4
	// We perform a bit-wise operator, knowing that number theory guarantees four perfect squares for n if this condition is met
	if n&uint64(7) == uint64(7) {
		return 4
	}
	remainder := n
	for remainder%4 == 0 {
		remainder /= 4
		if remainder&uint64(7) == uint64(7) {
			return 4
		}
	}

	// Step 3: check if n=2
	for i := uint64(1); i < root; i++ {
		i_2 := i * i
		a := uint64(math.Sqrt(float64(n - i_2)))
		if a*a+i_2 == n {
			return 2
		}
	}

	// Step 4: return the only remaining possibility, no need to prove it!
	return 3
}
