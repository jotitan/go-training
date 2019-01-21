package ex3_results

import (
	"fmt"
	"sync"
	"time"
)

//SumTwoFirstChannel sum only two first results in list
func SumTwoFirstChannel(channel chan int) int {
	sum := <-channel
	sum += <-channel
	return sum
}

//SumEverythingChannel sum everything in channel until close
func SumEverythingChannel(channel chan int) int {
	sum := 0
	for {
		if value, more := <-channel; more {
			sum += value
		} else {
			return sum
		}
	}
}

//ProduceSomething generates data in a channel every 100ms
func ProduceSomething(channel chan string) {
	counter := 0
	for {
		channel <- fmt.Sprintf("value : %d", counter)
		counter++
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

//ConsumeSomething read into channel and send, every 10 reading, results in other channel (like a list)
func ConsumeSomething(channel chan string, results chan []string) {
	outputList := make([]string, 0, 10)
	for {
		outputList = append(outputList, <-channel)
		if len(outputList) == 10 {
			results <- outputList
			outputList = make([]string, 0, 10)
		}
	}
}

//UseWaitGroup make nbToDone Done call on waitgroup
func UseWaitGroup(waiter *sync.WaitGroup, nbToDone int) {
	for i := 0; i < nbToDone; i++ {
		waiter.Done()
	}
}

//IsResponseOnTime return true if get a response before timeoutInMs
func IsResponseOnTime(response chan int, timeoutInMs int) bool {
	select {
	case <-response:
		return true
	case <-time.NewTimer(time.Duration(timeoutInMs) * time.Millisecond).C:
		return false
	}
}
