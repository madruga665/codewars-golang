package finduniq

import (
	"testing"
)

func TestFindUniq(t *testing.T) {
	want := 5
	result := FindUniq([]int{1, 1, 2, 2, 3, 3, 4, 4, 5})

	if want != result {
		t.Errorf("Want is %v but result is %v", want, result)
	}
}

func TestFindUniqNoUniq(t *testing.T) {
	want := 0
	result := FindUniq([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5})

	if want != result {
		t.Errorf("Want is %v but result is %v", want, result)
	}
}
