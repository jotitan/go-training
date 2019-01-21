package ex2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEx2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ex2 Suite")
}
