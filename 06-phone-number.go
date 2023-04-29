package kata

import (
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var sliceNumbers []string
	for _, v := range numbers {
		sliceNumbers = append(sliceNumbers, strconv.FormatUint(uint64(v), 10))
	}
	stringNumbers := strings.Join(sliceNumbers, "")
	return "(" + stringNumbers[0:3] + ") " + stringNumbers[3:6] + "-" + stringNumbers[6:10]
}
