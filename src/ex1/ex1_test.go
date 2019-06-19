package ex1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "ex1_results"
)

var _ = Describe("Ex1", func() {
	Describe("Test hello world", func() {
		Context("I call Hello world exemple", func() {
			It("Should return hello world", func() {
				Expect(HelloWorld()).To(Equal("Hello World"))
			})
		})
	})

	Describe("Mathematical operations", func() {
		Context("I add two numbers", func() {
			It("Should return 5 when I ask with 2 and 3", func() {
				Expect(AddInts(2, 3)).To(Equal(5))
			})
			It("Should return 4 when I ask with -12 and 16", func() {
				Expect(AddInts(-12, 16)).To(Equal(4))
			})
			It("Should return 26 when I ask with 11 as int64 and 32 as int15", func() {
				Expect(AddIntsDifferents(int64(11), int32(15))).To(Equal(26))
			})
		})
	})

	var superCali = "Supercalifragilisticexpialidocious"

	Describe("String operations", func() {
		Context("I substring a value", func() {
			It("Should return 'super' when I ask supercalifragilisticexpialidocious with 0,5", func() {
				Expect(Substring(superCali, 0, 5)).To(Equal("Super"))
			})
			It("Should return fragilistic when I ask supercalifragilisticexpialidocious with 9,20", func() {
				Expect(Substring(superCali, 9, 20)).To(Equal("fragilistic"))
			})
			It("Should return empty value when call with bad bounds 9,5", func() {
				Expect(Substring(superCali, 9, 5)).To(Equal(""))
			})
		})
		Context("I substring a value with error method", func() {
			It("Should return 'super' when I ask supercalifragilisticexpialidocious with 0,5", func() {
				value, err := SubstringWithErrors(superCali, 0, 5)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(value).Should(Equal("Super"))
			})
			It("Should return fragilistic when I ask supercalifragilisticexpialidocious with 9,20", func() {
				value, err := SubstringWithErrors(superCali, 9, 20)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(value).Should(Equal("fragilistic"))
			})
			It("Should throw error empty value when call with bad bounds 9,5", func() {
				value, err := SubstringWithErrors(superCali, 9, 5)
				Ω(err).Should(HaveOccurred())
				Ω(value).Should(Equal(""))
			})
		})
		Context("Extract numbers from string", func() {
			It("Sould found 0 values in 'Votre manque de foi me consterne'", func() {
				Expect(ExtractNumbersFromString("Votre manque de foi me consterne")).To(Equal([]int{}))
			})
			It("Sould found 3 values in 'Les valeurs 132 25, 602'", func() {
				Expect(ExtractNumbersFromString("Les valeurs 132 25, 602")).To(Equal([]int{132, 25, 602}))
			})
		})
	})

	Describe("Test swith types", func() {
		Context("Count types for ('val1',3,5,3.5,'val2',135,0.4)", func() {
			nbNumber, nbString, nbUnknown := CountTypes("val1", 3, 5, 3.5, "val2", 135, 0.4, make([]string, 0))
			It("Should find 2 strings", func() {
				Expect(nbString).To(Equal(2))
			})
			It("Should find 5 numbers", func() {
				Expect(nbNumber).To(Equal(5))
			})
			It("Should find 1 unknown", func() {
				Expect(nbUnknown).To(Equal(1))
			})
		})
	})

	Describe("List and map operations", func() {
		Context("Create set from values 13,54,65", func() {
			set := CreateSet(13, 54, 65)
			It("Should find 3 values in set", func() {
				Expect(len(set)).To(Equal(3))
			})
		})
	})
})
