package utils_test

import (
	. "github.com/YanshuoH/youkonger/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/YanshuoH/youkonger/consts"
)

var _ = Describe("Response", func() {
	Describe("NewJSONResponse", func() {
		Context("With optional argument as data", func() {
			It("Should return the argument value as data field", func() {
				By("string type as data")
				r := NewJSONResponse(consts.EventNotFound, "something wrong")
				Expect(r.ResultCode).To(Equal(consts.EventNotFound))
				Expect(r.ResultDescription).To(Equal(consts.Messenger.Get(consts.EventNotFound)))
				Expect(r.Data).To(Equal("something wrong"))

				By("array type as data")
				d1 := []string{"a", "b", "c"}
				r = NewJSONResponse(consts.OK, d1)
				Expect(r.Data).To(Equal(d1))

				By("struct type as data")
				d2 := struct {
					T int
				}{
					T: 123,
				}
				r = NewJSONResponse(consts.OK, d2)
				Expect(r.Data).To(Equal(d2))
			})
		})
	})

	Describe("NewOKJSONResponse", func() {
		Context("With provided data", func() {
			It("Should return a json response object with ok status", func() {
				d := struct {
					T string
				} {
					T: "yo",
				}
				r := NewOKJSONResponse(d)
				Expect(r.ResultCode).To(Equal(consts.OK))
				Expect(r.ResultDescription).To(Equal(consts.Messenger.Get(consts.OK)))
				Expect(r.Data).To(Equal(d))
			})
		})
	})
})
