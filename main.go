package main

import "FaisalBudiono/poc-go-concurrency/internal/scenario/scensync"

func main() {
	length := 1000

	scensync.Run(length)
}
