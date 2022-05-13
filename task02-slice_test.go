package homework

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {

	t.Run("Default_full", func(t *testing.T) {

		var input = []int64{1, 2, 5, 15}
		var result = []int64{15, 5, 2, 1}

		realResult := reverse(input)

		if !reflect.DeepEqual(realResult, result) {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}

	})

	t.Run("Default_nil ", func(t *testing.T) {

		var input []int64
		var result = []int64(nil)

		realResult := Reverse(input)

		if !reflect.DeepEqual(realResult, result) {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}

	})
}
