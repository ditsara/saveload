package saveload_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSaveload(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Saveload Suite")
}
