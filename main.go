package main

import (
	"fmt"

	"github.com/madruga665/codewars-golang/finduniq"
	"github.com/madruga665/codewars-golang/numbertostring"
	"github.com/madruga665/codewars-golang/positivesum"
	"github.com/madruga665/codewars-golang/spinwords"
	"github.com/madruga665/codewars-golang/summation"
)

func main() {
	resultSummation := summation.Summation(2)
	resultFindUniq := finduniq.FindUniq([]int{1, 1, 2, 2, 3, 3, 4, 4, 5})
	resultSpinWords := spinwords.SpinWords("testando o som")
	resultPositiveSum := positivesum.PositiveSum([]int{2, 5, -5, 6})
	resultNumberToString := numbertostring.NumberToString(665)

	fmt.Printf("Resultado so Summation: %d \n", resultSummation)
	fmt.Printf("Resultado so FindUniq: %d \n", resultFindUniq)
	fmt.Printf("Resultado so FindUniq: %s \n", resultSpinWords)
	fmt.Printf("Resultado so PositiveSum: %d \n", resultPositiveSum)
	fmt.Printf("Resultado so NumberToString: %s \n", resultNumberToString)
}
