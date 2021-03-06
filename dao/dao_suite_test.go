package dao_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/YanshuoH/youkonger/test"
	"testing"
)

func TestDao(t *testing.T) {
	RegisterFailHandler(Fail)

	// close when specs done
	defer RunSpecs(t, "Dao Suite")
}

var _ = BeforeSuite(func() {
	test.Setup()
})

var _ = AfterSuite(func() {
	test.Teardown()
})
