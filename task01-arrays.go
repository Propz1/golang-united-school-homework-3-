package homework

func Average(input [15]float32) (result float32) {

	var sum float32 = 0
	n := 0

	for _, i := range input {
		sum = sum + i
		if i != 0 {
			n++
		}
	}

	return sum / float32(n)
}
