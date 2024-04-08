package numbertostring

import "testing"

func TestNumberToString(t *testing.T) {
	want := "66"
	result := NumberToString(665)

	if want != result {
		t.Errorf("Expect result %s but received %s", want, result)
	}
}
