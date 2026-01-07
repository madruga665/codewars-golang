package pascalTriangle

import "fmt"

func PascalTriangle(number int) [][]int {
	triangle := [][]int{
		{1},
	}

	if number <= 0 {
		fmt.Println("O numero não pode ser igual ou menor a 0")

		return [][]int{}
	}

	for row := 1; row <= number; row++ {
		prevRow := triangle[row-1]
		newRow := []int{}

		for item := 0; item < len(prevRow)+1; item++ {
			var value1, value2 int

			if item-1 >= 0 {
				value1 = prevRow[item-1]
			} else {
				value1 = 0
			}

			if item < len(prevRow) {
				value2 = prevRow[item]
			} else {
				value2 = 0
			}

			newRow = append(newRow, value1+value2)

		}

		triangle = append(triangle, newRow)

	}

	return triangle
}
