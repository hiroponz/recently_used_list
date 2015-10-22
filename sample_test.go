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
			It("Countが0を返す", func() {
				Expect(sut.Count()).To(Equal(0))
			})
			It("Index('a')が-1を返す", func() {
				Expect(sut.Index("a")).To(Equal(-1))
			})
		})
		Context("文字列を'a'を追加する", func() {
			var sut *RecentlyUsedList
			BeforeEach(func() {
				sut = &RecentlyUsedList{}
				sut.Push("a")
			})
			It("IsEmptyがfalseを返す", func() {
				Expect(sut.IsEmpty()).To(BeFalse())
			})
			It("Countが1を返す", func() {
				Expect(sut.Count()).To(Equal(1))
			})
			It("Fetch(0)が追加した文字列を返す", func() {
				Expect(sut.Fetch(0)).To(Equal("a"))
			})
			It("Index('a')が0を返す", func() {
				Expect(sut.Index("a")).To(Equal(0))
			})
			It("Index('b')が-1を返す", func() {
				Expect(sut.Index("b")).To(Equal(-1))
			})
		})
		Context("文字列'a'->'b'を追加する", func() {
			var sut *RecentlyUsedList
			BeforeEach(func() {
				sut = &RecentlyUsedList{}
				sut.Push("a")
				sut.Push("b")
			})
			It("IsEmptyがfalseを返す", func() {
				Expect(sut.IsEmpty()).To(BeFalse())
			})
			It("Countが2を返す", func() {
				Expect(sut.Count()).To(Equal(2))
			})
			It("Fetch(0)が2番目に追加した文字列を返す", func() {
				Expect(sut.Fetch(0)).To(Equal("b"))
			})
			It("Fetch(1)が1番目に追加した文字列を返す", func() {
				Expect(sut.Fetch(1)).To(Equal("a"))
			})
			It("Index('a')が1を返す", func() {
				Expect(sut.Index("a")).To(Equal(1))
			})
			It("Index('b')が0を返す", func() {
				Expect(sut.Index("b")).To(Equal(0))
			})
			It("Index('c')が-1を返す", func() {
				Expect(sut.Index("c")).To(Equal(-1))
			})
		})
		Context("文字列'a'->'a'を追加する", func() {
			var sut *RecentlyUsedList
			BeforeEach(func() {
				sut = &RecentlyUsedList{}
				sut.Push("a")
				sut.Push("a")
			})
			It("IsEmptyがfalseを返す", func() {
				Expect(sut.IsEmpty()).To(BeFalse())
			})
			It("Countが1を返す", func() {
				Expect(sut.Count()).To(Equal(1))
			})
			It("Fetch(0)が2番目に追加した文字列を返す", func() {
				Expect(sut.Fetch(0)).To(Equal("a"))
			})
			It("Index('a')が0を返す", func() {
				Expect(sut.Index("a")).To(Equal(0))
			})
			It("Index('b')が-1を返す", func() {
				Expect(sut.Index("b")).To(Equal(-1))
			})
		})
	})
})
