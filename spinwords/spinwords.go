package spinwords

import "strings"

func SpinWords(phrase string) string {
	rune := []rune(phrase)
	inverted := []string{}

	for i := len(rune) - 1; i >= 0; i-- {
		inverted = append(inverted, string(rune[i]))
	}

	return strings.Join(inverted, "")
}
