package editor_test

import (
	"testing"

//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
	. "ginkgo"
	. "gomega")

func TestEditor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Editor Suite")
}
