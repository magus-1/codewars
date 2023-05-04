package kata

import "sort"

func SumOfIntervals(intervals [][2]int) int {
	// your code here
	var result int

	// Sort the intervals
	less := func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}
	sort.Slice(intervals, less)

	// Merge overlaps
	merged := [][2]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := len(merged) - 1
		prevInt, currInt := merged[last], intervals[i]

		if currInt[1] <= prevInt[1] {
			// full overlap, drop current interval
			continue
		}

		if currInt[0] <= prevInt[1] {
			// partial overlap, merge with previous interval
			last := len(merged) - 1
			merged[last][1] = currInt[1]
			continue
		}

		merged = append(merged, intervals[i])
	}

	// Measure the lengths
	for _, s := range merged {
		result += LengthInterval(s)
	}

	return result
}

func LengthInterval(interval [2]int) int {
	return interval[1] - interval[0]
}
