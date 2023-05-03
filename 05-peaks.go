package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	var isPeak bool = false
	var resultPos, resultPeaks []int
	var newPos, newPeak int

	for i, v := range array {
		if i == 0 || i == len(array)-1 {
			// skip beggining and end (cannot ever be peaks)
			continue
		}

		if isPeak == false && array[i-1] < v {
			// check if a peak is starting, register i, v
			isPeak = true
			newPos, newPeak = i, v
		}

		if isPeak == true && v < array[i+1] {
			// handles non-peaks outside peak start position
			isPeak = false
		}

		if isPeak == true && v > array[i+1] {
			// handles peaks register outside peak start position
			resultPos = append(resultPos, newPos)
			resultPeaks = append(resultPeaks, newPeak)
			isPeak = false
		}
	}
	return PosPeaks{resultPos, resultPeaks}
}
