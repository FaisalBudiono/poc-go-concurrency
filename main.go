package main

import (
	"FaisalBudiono/poc-go-concurrency/internal/scenario/scenworker"
)

func main() {
	length := 1000

	scenworker.Run(100, length)
}
