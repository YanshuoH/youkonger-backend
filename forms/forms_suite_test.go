package form_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/YanshuoH/youkonger/test"
	"testing"
)

func TestForms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forms Suite")
}

var _ = BeforeSuite(func() {
	test.Setup()
})

var _ = AfterSuite(func() {
	test.Teardown()
})
