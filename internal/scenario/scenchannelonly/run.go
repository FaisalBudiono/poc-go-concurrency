package scenchannelonly

import (
	"FaisalBudiono/poc-go-concurrency/internal/api"
	"FaisalBudiono/poc-go-concurrency/internal/gen"
)

func Run(length int) []int {
	inputs := gen.Input(length)

	totalInput := len(inputs)
	chanResults := make(chan int, totalInput)

	for _, input := range inputs {
		go worker(input, chanResults)
	}

	results := make([]int, len(inputs))
	for i := 0; i < totalInput; i++ {
		res := <-chanResults
		results[i] = res
	}

	return results
}

func worker(input int, res chan<- int) {
	r := api.Hit(input)

	res <- r
}
