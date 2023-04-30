package kata

// 04 kyu Pyramid Slide Down
func LongestSlideDown(pyramid [][]int) int {
	// Top-down approach would be a brute-force approach
	// Bottom-up approach: compute the sum for each pair, use this result to reduce the size of each layer
	// From O(2^n) complexity to O(n)
	bottom := len(pyramid) - 1
	for level := bottom; level > 0; level -= 1 {
		// sum every path for existing level
		currentSum := ReduceLevel(pyramid[level])

		// reduce total layers by 1
		for i := 0; i < len(pyramid[level])-1; i++ {
			pyramid[level-1][i] += currentSum[i]
		}
	}
	return pyramid[0][0]
}

func MaxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func ReduceLevel(levelSlice []int) []int {
	sumLevel := []int{}
	for i := 0; i < len(levelSlice)-1; i++ {
		sumLevel = append(sumLevel, MaxInt(levelSlice[i], levelSlice[i+1]))
	}
	return sumLevel
}
