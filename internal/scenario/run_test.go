package scenchannelonly_test

import (
	"FaisalBudiono/poc-go-concurrency/internal/scenario/scenchannelonly"
	"FaisalBudiono/poc-go-concurrency/internal/scenario/scenwaitgroup"
	"testing"
)

func BenchmarkChannelOnly_Run(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scenchannelonly.Run(i)
	}
}

func BenchmarkWaitGroup_Run(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scenwaitgroup.Run(i)
	}
}
