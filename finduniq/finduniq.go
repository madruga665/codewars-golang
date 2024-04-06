package finduniq

func FindUniq(numbers []int) int {
	countMap := map[int]int{}

	for _, num := range numbers {
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return 0
}
