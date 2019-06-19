package ex4_test

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "ex4"
)

var _ = Describe("Ex4", func() {
	Describe("Create light server", func() {
		BeforeSuite(func() {
			go RunServer(9007)
		})
		AfterSuite(func() {
			http.Get("http://localhost:9007/stop")
		})
		Context("Start a server with hello world", func() {
			It("Return hello world on /helloworld", func() {
				if resp, err := http.Get("http://localhost:9007/helloworld"); err == nil {
					data, _ := ioutil.ReadAll(resp.Body)
					Expect("Hello world").To(Equal(string(data)))
				} else {
					Fail("Impossible to request server")
				}
			})
		})
		Context("Call good url (add5), compare content", func() {
			It("I ask with 12 and want to get 17", func() {
				value, status := CallURL("http://localhost:9007/add5?value=12")
				Expect(200).To(Equal(status))
				Expect("17").To(Equal(value))
			})
			It("I ask with bad parameter and I get 403", func() {
				_, status := CallURL("http://localhost:9007/add5?value=toto")
				Expect(403).To(Equal(status))
			})
			It("I ask bad url and get 404", func() {
				_, status := CallURL("http://localhost:9007/bling")
				Expect(404).To(Equal(status))
			})
		})
		Context("Call get time and get it after 1 second", func() {
			It("I ask get time and want time", func() {
				begin := time.Now()
				_, status := CallURL("http://localhost:9007/getTime")
				end := time.Now()
				Expect(200).To(Equal(status))
				Expect(end.Sub(begin).Seconds()).To(BeNumerically(">=",1))
			})
			It("I ask 4 times getTime and wait more than 2 seconds from beginning", func() {
				begin := time.Now()
				nb := 4
				waiter := sync.WaitGroup{}
				waiter.Add(nb)
				for i := 0 ; i < nb ; i++ {
					go func(){
						CallURL("http://localhost:9007/getTime")
						waiter.Done()
					}()
				}
				waiter.Wait()
				end := time.Now()
				Expect(true).To(Equal(end.Sub(begin).Seconds() >= 2))
			})
		})
		Context("Call get time with limit many times and get 404", func() {
			It("I ask 4 times quikly and get at least one 404", func() {
				nb := 4
				response := make(chan int,nb)
				for i := 0 ; i < nb ; i++ {
					go func(){
						_, status := CallURL("http://localhost:9007/getTimeWithLimit")
						response <- status
					}()
				}
				sum := 0
				for i := 0 ; i < nb ; i++ {
					sum+=<-response
				}
				Expect(800).NotTo(Equal(sum))
			})
		})
		Context("Compute basket with slow resquest", func() {
			It("I ask price for one product, I get correct price and result between 200 and 250 ms", func() {
				begin := time.Now()
				data, status := CallURL("http://localhost:9007/computeBasket?products=1&quantities=1")
				requestTime := time.Now().Sub(begin).Nanoseconds()/1000000
				Expect(200).To(Equal(status))
				Expect(requestTime).To(And(BeNumerically(">=",200),BeNumerically("<",250)))
				Expect(float32(1.5)).To(Equal(parseFloat(data)))
			})
			It("I ask price for many products, I get correct price and result between 200 and 250 ms", func() {
				begin := time.Now()
				data, status := CallURL("http://localhost:9007/computeBasket?products=1,6,2,7,9&quantities=2,3,5,4,1")
				requestTime := time.Now().Sub(begin).Nanoseconds()/1000000
				Expect(200).To(Equal(status))
				Expect(requestTime).To(And(BeNumerically(">=",200),BeNumerically("<",250)))
				Expect(float32(94.4)).To(Equal(parseFloat(data)))
			})
			It("I ask price for all products, I get correct price and result between 200 and 250 ms", func() {
				begin := time.Now()
				data, status := CallURL("http://localhost:9007/computeBasket?products=1,2,3,4,5,6,7,8,9,10&quantities=1,1,1,1,1,1,1,1,1,1")
				requestTime := time.Now().Sub(begin).Nanoseconds()/1000000
				Expect(200).To(Equal(status))
				Expect(requestTime).To(And(BeNumerically(">=",200),BeNumerically("<",250)))
				Expect(float32(45.6)).To(Equal(parseFloat(data)))
			})
		})
	})
})

func parseFloat(value string)float32{
	if value,err := strconv.ParseFloat(value,32) ; err == nil {
		return float32(value)
	}
	return 0
}
