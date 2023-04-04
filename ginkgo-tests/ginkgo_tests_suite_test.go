package ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoTests Suite")
}
