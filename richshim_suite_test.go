package richshim_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRichshim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Richshim Suite")
}
