package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	result := []float64{}
	// copy signature
	for i, v := range signature {
		if i < n {
			result = append(result, v)
		}
	}
	// extend to n
	for i := 3; i < n; i++ {
		v := result[i-3] + result[i-2] + result[i-1]
		result = append(result, v)
	}
	return result
}
