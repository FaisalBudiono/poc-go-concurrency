package main

import "FaisalBudiono/poc-go-concurrency/internal/scenario/scenwaitgroup"

func main() {
	length := 1000

	scenwaitgroup.Run(length)
}
