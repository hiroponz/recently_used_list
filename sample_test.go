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
		Context("文字列を一つ追加する", func() {
			var sut *RecentlyUsedList
			BeforeEach(func() {
				sut = &RecentlyUsedList{}
				sut.Push("a")
			})
			It("IsEmptyがfalseを返す", func() {
				Expect(sut.IsEmpty()).To(BeFalse())
			})
			It("Fetchが追加した文字列を返す", func() {
				Expect(sut.Fetch(0)).To(Equal("a"))
			})
		})
	})
})
