package numbertostring

import "testing"

func TestNumberToString(t *testing.T) {
	want := "6655555"
	result := NumberToString(665)

	if want != result {
		t.Errorf("Expect result %s but received %s", want, result)
	}
}
