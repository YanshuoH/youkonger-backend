package utils_test

import (
	. "github.com/YanshuoH/youkonger/utils"

	"errors"
	"github.com/YanshuoH/youkonger/consts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errorwrap", func() {
	Describe("NewCommonError", func() {
		Context("With only code and nil as error", func() {
			It("Should return an error with code description as context", func() {
				cErr := NewCommonError("thing", nil)
				Expect(cErr.Code).Should(Equal("thing"))
				Expect(cErr.Description).Should(Equal(consts.Messenger.Get(consts.DefaultErrorMsg)))
				Expect(cErr.Detail).Should(Equal(""))
				Expect(cErr.Err.Error()).Should(Equal(consts.Messenger.Get(consts.DefaultErrorMsg)))
			})
		})

		Context("With defined description and detail", func() {
			It("Should use the defined arguments", func() {
				err := errors.New("yo")
				cErr := NewCommonError(consts.EventNotFound, errors.New("yo"), "description", "detail")
				Expect(cErr.Code).Should(Equal(consts.EventNotFound))
				Expect(cErr.Description).Should(Equal("description"))
				Expect(cErr.Detail).Should(Equal("yo"))
				// give it a wrap
				Expect(cErr.Err.Error()).Should(Equal(err.Error()))
			})
		})
	})
})
