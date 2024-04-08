package digitize

import (
	"strconv"
)

func Digitize(n int) []int {
	toText := strconv.Itoa(n)
	rune := []rune(toText)
	result := []int{}

	for i := len(rune) - 1; i >= 0; i-- {
		converted, _ := strconv.Atoi(string(rune[i]))
		result = append(result, converted)
	}

	return result
}
