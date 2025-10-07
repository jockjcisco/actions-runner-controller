package actions_test

import (
	"io"
	"os"
	"testing"
)

// testingTWrapper wraps *testing.T to implement ginkgo.GinkgoTInterface
type testingTWrapper struct {
	*testing.T
}

// Attr implements the missing method from ginkgo.GinkgoTInterface
func (w *testingTWrapper) Attr(key, value string) {
	// This is a no-op for testing.T compatibility
}

// Output implements the missing method from ginkgo.GinkgoTInterface  
func (w *testingTWrapper) Output() io.Writer {
	// Return a basic writer for testing.T compatibility
	return os.Stdout
}
