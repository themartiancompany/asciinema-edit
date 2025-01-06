package commands_test

import (
	"testing"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
	. "ginkgo"
	. "gomega"

)

func TestCommands(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands Suite")
}
