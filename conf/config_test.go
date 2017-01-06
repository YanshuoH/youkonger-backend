package conf_test

import (
	. "github.com/YanshuoH/youkonger/conf"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Describe("Setup", func() {
		Context("With un-recognized file path", func() {
			It("Should return an error", func() {
				c, err := Setup("")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("is a directory"))
				Expect(c).To(BeNil())
			})
		})

		Context("With un-readable file content", func() {
			It("Should return an error", func() {
				c, err := Setup("conf_mock.toml")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Near line"))
				Expect(c).To(BeNil())
			})
		})

		Context("With correct file", func() {
			It("Should return the right conf", func() {
				c, err := Setup("conf_loc.toml")
				Expect(err).NotTo(HaveOccurred())
				Expect(c.AppConf.GinMode).To(Equal(gin.DebugMode))
			})
		})
	})
})
