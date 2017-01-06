package consts_test

import (
	. "github.com/YanshuoH/youkonger/consts"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Codes", func() {
	Describe("Messenger", func() {
		Describe("Get", func() {
			Context("With defined error code", func() {
				It("Should return the mapped message", func() {
					Expect(Messenger.Get(FormSaveError)).To(Equal("保存时发生了意外错误, 请稍候重试"))
				})
			})

			Context("With unknown error code", func() {
				It("Should return the default message", func() {
					Expect(Messenger.Get("something not defined")).To(Equal(Messenger.Get(DefaultErrorMsg)))
				})
			})
		})
	})
})
