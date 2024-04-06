package numbertostring

import "strconv"

func NumberToString(number int) string {
	converted := strconv.Itoa(number)

	return converted
}
