package rpctest

import (
	"testing"
	"os"
)

// Needs to be in a _test.go file to be picked up
func TestMain(m *testing.M) {
	returnValue := TestWrapper(func() int {
		return m.Run()
	})

	defer os.Exit(returnValue)
}

