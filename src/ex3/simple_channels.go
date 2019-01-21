package ex3

import "sync"

//SumTwoFirstChannel sum only two first results in list
func SumTwoFirstChannel(channel chan int) int {
	return 0
}

//SumEverythingChannel sum everything in channel until close
func SumEverythingChannel(channel chan int) int {
	return 0
}

//ProduceSomething generates data in a channel every 100ms
func ProduceSomething(channel chan string) {}

//ConsumeSomething read into channel and send, every 10 reading, results in other channel (like a list)
func ConsumeSomething(channel chan string, results chan []string) {}

//IsResponseOnTime return true if get a response before timeoutInMs
func IsResponseOnTime(response chan int, timeoutInMs int) bool {
	return false
}

//UseWaitGroup make nbToDone Done call on waitgroup
func UseWaitGroup(waiter *sync.WaitGroup, nbToDone int) {}
