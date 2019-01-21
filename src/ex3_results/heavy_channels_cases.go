package ex3_results

import (
	"sync"
	"time"
)

/* Exercices sur les channels */

func fctToRun(val int) int {
	time.Sleep(time.Duration(1) * time.Second)
	return val * 2
}

func makeSum(sum *int, chanel chan int, waiter *sync.WaitGroup) {
	for {
		*sum += <-chanel
		waiter.Done()
	}
}

func makeSum2(chanel chan int, waiter *sync.WaitGroup, chanelOut chan int) {
	sum := 0
	for {
		if v, more := <-chanel; more {
			sum += v
			waiter.Done()
		} else {
			break
		}
	}
	chanelOut <- sum
}

// ComputeSum fourth version of compute, go routine, waiter and chanel and chanel to get sub sum
func ComputeSum(nbCycles int) (int, int) {
	waiter := sync.WaitGroup{}
	waiter.Add(nbCycles)
	chanel := make(chan int, 100)
	sumChanel := make(chan int, 1)
	nbThreads := 3
	for i := 0; i < nbThreads; i++ {
		go makeSum2(chanel, &waiter, sumChanel)
	}
	// Limit number of goroutines in // => Create too many and take too many memory
	limiter := make(chan struct{}, 10000)
	for i := 0; i < nbCycles; i++ {
		limiter <- struct{}{}
		go func(v int) {
			chanel <- fctToRun(v)
		}(i)
		<-limiter
	}
	waiter.Wait()
	close(chanel)

	sum := 0
	for i := 0; i < nbThreads; i++ {
		sum += <-sumChanel
	}
	return sum, nbCycles * (nbCycles - 1)
}

// ComputeSum3 third version of compute, go routine, waiter and chanel
func ComputeSum3(nbCycles int) (int, int) {
	sumToHave := nbCycles * (nbCycles - 1)
	sum := 0
	waiter := sync.WaitGroup{}
	waiter.Add(nbCycles)
	chanel := make(chan int, 100)
	go makeSum(&sum, chanel, &waiter)

	for i := 0; i < nbCycles; i++ {
		go func(v int) {
			chanel <- fctToRun(v)
		}(i)
	}
	waiter.Wait()
	return sum, sumToHave
}

// ComputeSumV2 thrid version of compute, go routine and waiter (to wait the end)
func ComputeSumV2(nbCycles int) (int, int) {
	sumToHave := nbCycles * (nbCycles - 1)
	sum := 0
	waiter := sync.WaitGroup{}
	waiter.Add(nbCycles)
	for i := 0; i < nbCycles; i++ {
		go func(v int) {
			sum += fctToRun(v)
			waiter.Done()
		}(i)
	}
	waiter.Wait()
	return sum, sumToHave
}

// ComputeSumV1 second version of compute, go routine only
func ComputeSumV1(nbCycles int) (int, int) {
	sumToHave := nbCycles * (nbCycles - 1)
	sum := 0
	for i := 0; i < nbCycles; i++ {
		go func(v int) {
			sum += fctToRun(v)
		}(i)
	}
	return sum, sumToHave
}

// ComputeSumV0 first version of compute, no goroutine
func ComputeSumV0(nbCycles int) (int, int) {
	sumToHave := nbCycles * (nbCycles - 1)
	sum := 0
	for i := 0; i < nbCycles; i++ {
		sum += fctToRun(i)
	}
	return sum, sumToHave
}
