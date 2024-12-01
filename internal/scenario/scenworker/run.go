package scenworker

import (
	"FaisalBudiono/poc-go-concurrency/internal/api"
	"FaisalBudiono/poc-go-concurrency/internal/gen"
	"fmt"
	"time"
)

type response struct {
	input int
	res   int
}

func Run(workerNumber, length int) []int {
	start := time.Now()

	inputs := gen.Input(length)

	totalInput := len(inputs)
	jobs := make(chan int, totalInput)
	chanResults := make(chan response, totalInput)

	for i := 0; i < workerNumber; i++ {
		go worker(jobs, chanResults)
	}

	for _, input := range inputs {
		fmt.Printf("Sending job %d\n", input)
		jobs <- input
	}
	close(jobs)

	results := make([]int, len(inputs))
	for i := 0; i < totalInput; i++ {
		res := <-chanResults
		fmt.Printf("Received result %d\n", res.input)
		results[i] = res.res
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)

	return results
}

func worker(job <-chan int, res chan<- response) {
	for input := range job {
		fmt.Printf("Working job %d\n", input)
		r := api.Hit(input)
		res <- response{input, r}
	}
}
