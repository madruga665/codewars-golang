package summation

func Summation(number int) int {
	var numbers = []int{}
	var result int

	for i := 0; i <= number; i++ {
		numbers = append(numbers, i)
	}

	for _, n := range numbers {
		result += n
	}

	return result
}
