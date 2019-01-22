package ex4_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "ex4_results"
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
	})
})
