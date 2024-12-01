package scenwaitgroup

import (
	"FaisalBudiono/poc-go-concurrency/internal/api"
	"FaisalBudiono/poc-go-concurrency/internal/gen"
	"sync"
)

func Run(length int) []int {
	inputs := gen.Input(length)

	totalInput := len(inputs)
	chanResults := make(chan int, totalInput)

	var wg sync.WaitGroup
	wg.Add(totalInput)

	for _, input := range inputs {
		go worker(input, chanResults, &wg)
	}

	wg.Wait()

	results := make([]int, len(inputs))
	for i := 0; i < totalInput; i++ {
		res := <-chanResults
		results[i] = res
	}
	close(chanResults)

	return results
}

func worker(input int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	r := api.Hit(input)
	res <- r
}
