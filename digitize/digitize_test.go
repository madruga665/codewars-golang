package digitize

import (
	"reflect"
	"testing"
)

func TestDigitize(t *testing.T) {
	want := []int{1, 3, 2, 5, 3}
	result := Digitize(35231)

	if !reflect.DeepEqual(want, result) {
		t.Errorf("Expect result %v but received %v", want, result)
	}
}
