package jrenders_test

import (
	. "github.com/YanshuoH/youkonger/jrenders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventDate", func() {

	Describe("Itemize", func() {
		Context("With provided eventDate and event", func() {
			It("Should return the struct for json rendering", func() {
				ed := eventDateSet[0]
				e := eventSet[0]
				j := EventDate.Itemize(&ed, &e)
				Expect(j.TimeInUnix).To(Equal(ed.Time.Unix()))
				Expect(j.JEventParticipants.JList).To(HaveLen(1))
			})
		})
	})

	Describe("List", func() {
		Context("With provided eventDates and event", func() {
			It("Should return the struct for json rendering", func() {
				e := eventSet[0]
				eds := eventDateSet
				j := EventDate.List(eds, &e)
				Expect(j.JList).To(HaveLen(len(eds)))
			})
		})
	})
})
