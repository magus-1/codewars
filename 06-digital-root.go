package kata

import (
	"strconv"
	"strings"
)

func DigitalRoot(n int) int {
	m := strconv.Itoa(n)
	result, _ := strconv.Atoi(FindRoot(m))

	return result
}

func FindRoot(m string) string {
	var n int
	var result string
	for _, v := range strings.Split(m, "") {
		nAdd, _ := strconv.Atoi(v)
		n += nAdd
	}
	result = strconv.Itoa(n)
	if n >= 10 {
		result = FindRoot(result)
	}
	return result
}
