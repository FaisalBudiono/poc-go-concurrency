package api

import (
	"math"
	"time"
)

func Hit(id int) int {
	time.Sleep(time.Microsecond * 532)

	return int(math.Pow(float64(id), 3))
}
