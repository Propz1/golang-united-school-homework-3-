package homework

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {

	t.Run("Default_full", func(t *testing.T) {

		var input = []int64{1, 2, 5, 15}
		var result = []int64{15, 5, 2, 1}

		realResult := Reverse(input)

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

func TestSortMapValues(t *testing.T) {

	t.Run("One", func(t *testing.T) {

		var input = map[int]string{2: "a", 0: "b", 1: "c"}
		var result = []string{"b", "c", "a"}

		realResult := SortMapValues(input)

		if !reflect.DeepEqual(realResult, result) {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}

	})

	t.Run("Two", func(t *testing.T) {

		var input = map[int]string{10: "aa", 0: "bb", 500: "cc"}
		var result = []string{"bb", "aa", "cc"}

		realResult := SortMapValues(input)

		if !reflect.DeepEqual(realResult, result) {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}
	})

	t.Run("three", func(t *testing.T) {

		var result []string
		var input = map[int]string{}

		realResult := SortMapValues(input)

		if !reflect.DeepEqual(realResult, result) {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}
	})

}
