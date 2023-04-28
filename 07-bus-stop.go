package kata

func Number(busStops [][2]int) int {
	result := 0
	for _, v := range busStops {
		result += v[0] - v[1]
	}
	return result
}
