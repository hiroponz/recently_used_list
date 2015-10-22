package go_ginkgo_test

import (
	. "github.com/tddbc/go_ginkgo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
	Describe("#RecentlyUsedList", func() {
		Context("リストが空の時", func() {
			var sut *RecentlyUsedList
			BeforeEach(func() {
				sut = &RecentlyUsedList{}
			})
			It("IsEmptyがtrueを返す", func() {
				Expect(sut.IsEmpty()).To(BeTrue())
			})
		})
	})
})
