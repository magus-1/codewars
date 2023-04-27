package kata

import (
	"math/big"
)

func IsPrime(n int) bool {
	switch {
	case n < 2:
		return false
	default:
		target := big.NewInt(int64(n))
		return target.ProbablyPrime(5)
	}
}
