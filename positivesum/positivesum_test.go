package positivesum

import "testing"

func TestPositiveSum(t *testing.T) {
	want := 13
	result := PositiveSum([]int{2, 5, -5, 6})

	if want != result {
		t.Errorf("Expect result is %d but received %d", want, result)
	}
}
