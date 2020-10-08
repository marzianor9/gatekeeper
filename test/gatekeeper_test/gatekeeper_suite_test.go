package gatekeeper_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGatekeeperTest(t *testing.T) {
	fmt.Println("i am suite")
	RegisterFailHandler(Fail)
	RunSpecs(t, "GatekeeperTest Suite")
}
