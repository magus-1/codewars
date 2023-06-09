package kata

import (
	"fmt"
	"strings"
)

func DecodeMorse(morseCode string) string {
	newCode := strings.TrimSpace(morseCode)
	newCode = strings.ReplaceAll(newCode, "   ", " ....... ")
	listCodes := strings.Split(newCode, " ")
	var result string = ""
	for _, v := range listCodes {
		if v == "......." {
			result += " "
		} else {
			// result += MORSE_CODE[v]
			result += v // replacement line to avoid error in local repo
		}
	}
	fmt.Println("Result:", result, "\nmorseCode:", morseCode, "\nlistCodes:", listCodes)
	return result
}
