package ex3

import "time"

func fctToRun(val int) int {
	time.Sleep(time.Duration(1) * time.Second)
	return val * 2
}

// ComputeSum improve method to answer as fast as possible
func ComputeSum(nbCycles int) (int, int) {
	sum := 0
	for i := 0; i < nbCycles; i++ {
		sum += fctToRun(i)
	}
	return sum, nbCycles * (nbCycles - 1)
}
