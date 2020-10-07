package gatekeeper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGatekeeperTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GatekeeperTest Suite")
}
