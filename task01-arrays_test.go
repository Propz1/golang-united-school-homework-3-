package homework

import (
	"testing"
)

func TestAverage(t *testing.T) {

	var result float32
	var input = [15]float32{1, 2, 3, 4, 5, 6}

	result = 3.5
	realResult := average(input)

	if result != realResult {
		t.Errorf("expected result %f != %f real result", result, realResult)
	}

}
