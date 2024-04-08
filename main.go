package main

import (
	"fmt"

	"github.com/madruga665/codewars-golang/digitize"
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
	resultDigitize := digitize.Digitize(665)

	fmt.Printf("Resultado do Summation: %d \n", resultSummation)
	fmt.Printf("Resultado do FindUniq: %d \n", resultFindUniq)
	fmt.Printf("Resultado do FindUniq: %s \n", resultSpinWords)
	fmt.Printf("Resultado do PositiveSum: %d \n", resultPositiveSum)
	fmt.Printf("Resultado do NumberToString: %s \n", resultNumberToString)
	fmt.Printf("Resultado do resultDigitize: %d \n", resultDigitize)
}
