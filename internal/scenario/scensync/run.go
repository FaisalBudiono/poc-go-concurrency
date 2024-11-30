package scensync

import (
	"FaisalBudiono/poc-go-concurrency/internal/api"
	"FaisalBudiono/poc-go-concurrency/internal/gen"
	"fmt"
	"time"
)

func Run(length int) {
	start := time.Now()

	inputs := gen.Input(length)
	results := make([]int, len(inputs))

	for i, input := range inputs {
		results[i] = api.Hit(input)

		fmt.Printf(
			"Hitting %d finish on %s\n",
			input,
			time.Now().Sub(start),
		)
	}

	finishOn := time.Now().Sub(start)
	fmt.Printf("Finish on %s\n\n", finishOn)

	fmt.Printf("Result %#v\n\n", results)
}
