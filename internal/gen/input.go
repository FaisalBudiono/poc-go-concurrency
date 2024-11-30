package gen

func Input(length int) []int {
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		inputs[i] = i
	}

	return inputs
}
