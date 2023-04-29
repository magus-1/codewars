package kata

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	// input:
	// 	numbers: slice of ints
	// 	target: int
	// output:
	// 	result: array of ints, pair of indeces holding location of a pair of ints whose sum equals target
	// the numbers slice cannot be sorted as the order matters for result. So I will not try to optimize brute force
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if target == numbers[i]+numbers[j] {
				return [2]int{i, j}
			}
		}
	}
	fmt.Println("No solution found")
	return [2]int{0, 0}
}
