package ex4_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "ex4"
)

var _ = Describe("Ex4", func() {
	Describe("Create light server", func() {
		Context("Start a server with hello world", func() {
			BeforeEach(func() {
				go RunServer(9007)
			})
			AfterEach(func() {
				http.Get("http://localhost:9007/stop")
			})
			It("Return hello world on /helloworld", func() {
				if resp, err := http.Get("http://localhost:9007/helloworld"); err == nil {
					data, _ := ioutil.ReadAll(resp.Body)
					Expect("Hello world").To(Equal(string(data)))
				} else {
					Fail("Impossible to request server")
				}
			})
		})
	})
})
