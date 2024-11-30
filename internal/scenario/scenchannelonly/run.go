package scenchannelonly

import (
	"FaisalBudiono/poc-go-concurrency/internal/api"
	"FaisalBudiono/poc-go-concurrency/internal/gen"
	"fmt"
	"time"
)

func Run(length int) []int {
	var start time.Time = time.Now()

	inputs := gen.Input(length)

	totalInput := len(inputs)
	chanResults := make(chan int, totalInput)

	for _, input := range inputs {
		fmt.Printf("Sending %d\n\n", input)
		go worker(start, input, chanResults)
	}

	results := make([]int, len(inputs))
	for i := 0; i < totalInput; i++ {
		res := <-chanResults
		results[i] = res
	}

	finishOn := time.Now().Sub(start)
	fmt.Printf("Finish on %s\n\n", finishOn)

	return results
}

func worker(n time.Time, input int, res chan<- int) {
	r := api.Hit(input)
	fmt.Printf(
		"Hitting %d finish on %s\n\n",
		input,
		time.Now().Sub(n),
	)

	res <- r
}
