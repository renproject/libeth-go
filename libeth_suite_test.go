package libeth_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLibETH(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LibETH Test Suite")
}
