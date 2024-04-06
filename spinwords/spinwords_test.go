package spinwords

import "testing"

func TestSpinWords(t *testing.T) {
	want := "mos o odnatset"
	result := SpinWords("testando o som")

	if want != result {
		t.Errorf("Want is %s but result is %s", want, result)
	}
}
