package ex3_test

import (
	. "ex3"
	"fmt"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ex3", func() {
	Describe("Use channel for simple cases", func() {
		Context("Compute sum in channel", func() {
			channelSum := make(chan int, 4)

			channelSum <- 2
			channelSum <- 8
			It("Must get 10 as sum", func() {
				Expect(10).To(Equal(SumTwoFirstChannel(channelSum)))
			})
			channelSum <- 14
			channelSum <- 25
			It("Must get 39 as sum", func() {
				Expect(39).To(Equal(SumTwoFirstChannel(channelSum)))
			})
		})
		Context("Compute everything in channel", func() {
			channelSum := make(chan int, 10)
			for i := 1; i <= 10; i++ {
				channelSum <- i
			}
			close(channelSum)
			It("Total sum must be 55", func() {
				Expect(55).To(Equal(SumEverythingChannel(channelSum)))
			})
		})
		Context("Producer / consumer cases", func() {
			mainChannel := make(chan string, 100)
			go ProduceSomething(mainChannel)

			outputResults := make(chan []string, 10)
			go ConsumeSomething(mainChannel, outputResults)

			// Read some elements
			It("Read some elements in channel output", func() {
				for i := 0; i < 3; i++ {
					select {
					case list := <-outputResults:
						Expect(10).To(Equal(len(list)))
					case <-time.NewTimer(time.Duration(2) * time.Second).C:
						Fail("Timeout before getting results")
					}
				}
			})
		})
		Context("Use timer to detect timeout", func() {
			channel := make(chan int, 1)
			BeforeEach(func() {
				go func() {
					time.Sleep(time.Duration(1) * time.Second)
					channel <- 1
				}()
			})
			It("Must get timeout when wait too long", func() {
				Expect(false).To(Equal(IsResponseOnTime(channel, 500)))
			})
			It("Must get no timeout", func() {
				Expect(true).To(Equal(IsResponseOnTime(channel, 2000)))
			})
		})
		Context("Use wait group", func() {
			channelEnd := make(chan struct{}, 1)
			BeforeEach(func() {
				waiter := sync.WaitGroup{}
				waiter.Add(15)
				go UseWaitGroup(&waiter, 15)

				go func() {
					waiter.Wait()
					channelEnd <- struct{}{}
				}()
			})
			// Wait 1 second to get result, force fail
			It("Should get notification for end waiter instead of timeout", func() {
				select {
				case <-time.NewTimer(time.Duration(1) * time.Second).C:
					Expect(true).To(Equal(false))
				case <-channelEnd:
					Expect(true).To(Equal(true))
				}
			})
		})
	})

	Describe("Use channel to compute sum as fast as possible", func() {
		Context("Run 10 cycles (increase number after)", func() {
			sum, goodSum := ComputeSum(10)
			It(fmt.Sprintf("Must get %d", goodSum), func() {
				Expect(sum).To(Equal(goodSum))
			})
		})
	})
})
