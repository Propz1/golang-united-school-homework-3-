package homework

func average(input [15]float32) (result float32) {

	var sum float32 = 0

	for _, i := range input {
		sum = sum + i
	}

	return sum / float32(len(input))
}
