package scensync

import (
	"FaisalBudiono/poc-go-concurrency/internal/api"
	"FaisalBudiono/poc-go-concurrency/internal/gen"
)

func Run(length int) []int {
	inputs := gen.Input(length)
	results := make([]int, len(inputs))

	for i, input := range inputs {
		results[i] = api.Hit(input)
	}

	return results
}
