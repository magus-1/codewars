package main

import (
	"fmt"
	"math/big"
	"time"
)

func IsPrime(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	default:
		target := big.NewInt(int64(n))
		return target.ProbablyPrime(5)
	}
	// not a base case and no valid divisor found
	return true
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	fmt.Println(
		IsPrime(10000),
		IsPrime(100000),
		IsPrime(1000000))
}
