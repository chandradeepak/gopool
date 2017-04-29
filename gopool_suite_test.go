package gopool_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGopool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gopool Suite")
}
