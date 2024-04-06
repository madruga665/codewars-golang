package summation

import "testing"

func TestSummation(t *testing.T) {
	want := 3
	result := Summation(2)

	if want != result {
		t.Errorf("Want is %d but result is %d", want, result)
	}

}
