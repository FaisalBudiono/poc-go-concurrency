package scenwaitgroup

import (
	"FaisalBudiono/poc-go-concurrency/internal/api"
	"FaisalBudiono/poc-go-concurrency/internal/gen"
	"fmt"
	"sync"
	"time"
)

func Run(length int) {
	start := time.Now()

	inputs := gen.Input(length)

	totalInput := len(inputs)
	chanResults := make(chan int, totalInput)

	var wg sync.WaitGroup
	wg.Add(totalInput)

	for _, input := range inputs {
		go worker(start, input, chanResults, &wg)
	}

	wg.Wait()

	results := make([]int, len(inputs))
	for i := 0; i < totalInput; i++ {
		res := <-chanResults
		results[i] = res
	}
	close(chanResults)

	finishOn := time.Now().Sub(start)
	fmt.Printf("Finish on %s\n\n", finishOn)

	fmt.Printf("Result(%d) %#v\n\n", len(results), results)
}

func worker(n time.Time, input int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	r := api.Hit(input)
	fmt.Printf(
		"Hitting %d finish on %s\n",
		input,
		time.Now().Sub(n),
	)

	res <- r
}
