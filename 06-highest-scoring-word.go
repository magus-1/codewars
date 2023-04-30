package kata

import (
	"fmt"
	"strings"
)

// 06 kyu Highest Scoring Word

var dictSlice = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func High(s string) string {
	// input: s string, a set of words
	// output: string, the word with the highest value
	// value is obtained by the sum of each letter (a=1, b=2, etc)
	var maxWord string
	currentValue, maxValue := 0, 0
	wordSlice := strings.Split(s, " ")
	for _, v := range wordSlice {
		currentValue = Value(v)
		if currentValue > maxValue {
			maxValue = currentValue
			maxWord = v
		}
	}
	return maxWord
}

func Value(w string) int {
	var wordValue = 0
	l := strings.Split(w, "")
	for _, v := range l {
		wordValue += FindValue(v)
	}
	return wordValue
}

func FindValue(s string) int {
	for i, v := range dictSlice {
		if v == s {
			return i + 1
		}
	}
	fmt.Println("Letter not found!")
	return 0
}
