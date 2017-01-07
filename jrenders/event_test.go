package jrenders_test

import (
	. "github.com/YanshuoH/youkonger/jrenders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {
	Describe("Itemize", func() {
		Context("With provided event and param", func() {
			It("Should return the struct for json rendering with specified fields", func() {
				e := eventSet[0]
				By("No admin hash")
				j := Event.Itemize(&e, EventParam{
					ShowHash: false,
				})
				Expect(j.AdminHash).To(Equal(""))
				Expect(j.JEventDates.JList).To(HaveLen(1))
			})
		})
	})
})
