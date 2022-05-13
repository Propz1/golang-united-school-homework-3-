package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {

	if len(input) != 0 {

		keys := make([]int, 0, len(input))
		for k := range input {
			keys = append(keys, k)
		}

		sort.Ints(keys)

		for _, k := range keys {
			result = append(result, input[k])
		}

	}

	return result
}
