package ex4_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEx4(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ex4 Suite")
}
