package gatekeeper_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/open-policy-agent/gatekeeper/test/gatekeeper_test/utils"
)

var _ = Describe("Gatekeeper", func() {
	fmt.Println("i am test")

	It("should confirm basic math", func() {
		Expect(1).To(Equal(2 / 2))
	})

	It("should confirm namespace name", func() {
		Expect(utils.TestNamespace).To(Equal("gatekeeper_test"))
	})
})
