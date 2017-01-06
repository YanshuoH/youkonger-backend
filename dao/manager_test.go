package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	Describe("GetManager", func() {
		Context("With given tx", func() {
			It("Should return the given tx as db inside manager", func() {
				tx := Conn.Begin()
				em := GetManager(tx)
				Expect(&em.DB).To(Equal(&tx))
			})
		})

		Context("With empty arg", func() {
			It("Should return an address other than the tx's address", func() {
				tx := Conn.Begin()
				em := GetManager()
				Expect(&em.DB).NotTo(Equal(&tx))
			})
		})
	})

	Describe("Event", func() {
		It("Should return an Event dao", func() {
			em := GetManager()
			Expect(em.Event()).NotTo(BeNil())
		})
	})

	Describe("EventDate", func() {
		It("Should return an EventDate dao", func() {
			em := GetManager()
			Expect(em.EventDate()).NotTo(BeNil())
		})
	})

	Describe("EventParticipant", func() {
		It("Should return an EventParticipant dao", func() {
			em := GetManager()
			Expect(em.EventParticipant()).NotTo(BeNil())
		})
	})
})
