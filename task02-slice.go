package homework

func reverse(input []int64) (result []int64) {

	if input != nil {

		i := len(input) - 1

		result = make([]int64, i+1)
		copy(result, input)
		j := 0
		for i >= 0 {
			result[j] = input[i]
			i--
			j++
		}
	}

	return result
}
