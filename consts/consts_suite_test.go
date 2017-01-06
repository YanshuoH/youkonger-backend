package consts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConsts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Consts Suite")
}
