package main

import (
	"FaisalBudiono/poc-go-concurrency/internal/scenario/scenchannelonly"
)

func main() {
	length := 1000

	scenchannelonly.Run(length)
}
