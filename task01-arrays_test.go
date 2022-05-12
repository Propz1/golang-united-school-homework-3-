package homework

import (
	"testing"
)

func TestAverage(t *testing.T) {

	var input = [15]float32{1, 2, 3, 4, 5, 6}
	var result float32

	t.Run("Array", func(t *testing.T) {

		result = 3.5
		realResult := Average(input)

		if result != realResult {
			t.Errorf("expected result %f != %f real result", result, realResult)
		}
	})

}
