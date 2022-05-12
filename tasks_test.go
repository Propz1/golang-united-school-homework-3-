package homework

import (
	"reflect"
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

func TestReverse(t *testing.T) {

	var input = []int64{1, 2, 5, 15}
	var result = []int64{15, 5, 2, 1}

	realResult := Reverse(input)

	if !reflect.DeepEqual(realResult, result) {
		t.Errorf("expected result %v != %v real result", result, realResult)
	}

}
