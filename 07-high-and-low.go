package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	// Code here
	strSlice := strings.Split(in, " ")
	var intSlice []int64
	var pInt int64

	for _, s := range strSlice {
		pInt, _ = strconv.ParseInt(s, 10, 64)
		intSlice = append(intSlice, pInt)
	}
	low, high := MinMax(intSlice)
	return strconv.FormatInt(high, 10) + " " + strconv.FormatInt(low, 10)
}

func MinMax(intSlice []int64) (int64, int64) {
	var max int64 = intSlice[0]
	var min int64 = intSlice[0]
	for _, value := range intSlice {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	fmt.Println(strconv.ParseInt("-1", 10, 64))
	fmt.Println(HighAndLow("1 2 3 4 5"))  // return "5 1"
	fmt.Println(HighAndLow("1 2 -3 4 5")) // return "5 -3"
	fmt.Println(HighAndLow("1 9 3 4 -5")) // return "9 -5"
}
