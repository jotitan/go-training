package ex3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEx3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ex3 Suite")
}
