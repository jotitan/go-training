package ex1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEx1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ex1 Suite")

}
